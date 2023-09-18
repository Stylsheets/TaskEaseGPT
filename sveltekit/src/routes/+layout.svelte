<script lang="ts">
    import "../app.css";
    import Sidebar from "../components/Sidebar.svelte";
    import Button from "../components/Button.svelte";
    import {slide} from "svelte/transition";
    import {isSidebarOpen} from "../writeables";
    import {cubicOut} from "svelte/easing";


    function updateButtonData() {
        const sidebarButton = document.querySelector('.sidebar-button');
        if (!sidebarButton) return;
        sidebarButton.setAttribute('data-active', String($isSidebarOpen));
    }

    function toggleSidebar() {
        isSidebarOpen.update((value) => !value);
        updateButtonData();
    }

    let screenWidth: number;
    $: {
        if (screenWidth > 1024) {
            isSidebarOpen.set(true);
            updateButtonData();
        }
    }
</script>

<svelte:window bind:innerWidth={screenWidth} />

<div data-active="true" class="sidebar-button">
<Button func={toggleSidebar} class="max-w-fit p-3">
    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
        <path fill-rule="evenodd"
              d="M3 6.75A.75.75 0 013.75 6h16.5a.75.75 0 010 1.5H3.75A.75.75 0 013 6.75zM3 12a.75.75 0 01.75-.75h16.5a.75.75 0 010 1.5H3.75A.75.75 0 013 12zm0 5.25a.75.75 0 01.75-.75H12a.75.75 0 010 1.5H3.75a.75.75 0 01-.75-.75z"
              clip-rule="evenodd"/>
    </svg>
</Button>
</div>
<div class="grid-container {$isSidebarOpen ? '' : 'sidebar-closed'}">
    {#if $isSidebarOpen}
        <aside transition:slide={
        {axis: 'x', duration: 300, easing: cubicOut}
        } class="sidebar">
            <Sidebar/>
        </aside>
    {/if}

    <main class="content">
        <slot/>
    </main>
</div>

<style lang="postcss">
    @media screen and (min-width: 1024px) {
        .sidebar-button {
            display: none;
        }
    }

    .sidebar-button {
        position: absolute;
        top: 5px;
        right: calc(100vw - 54px);
        transition: all 0.3s ease-in-out;
        z-index: 10;
    }

    .sidebar-button[data-active="true"] {
        right: 5px;
    }

    .grid-container {
        display: grid;
        grid-template-rows: 100%;
        grid-template-columns: 300px 1fr;
        grid-template-areas: "sidebar content";
        height: 100vh;
    }

    .sidebar-closed {
        grid-template-columns: 0px 1fr;
    }

    .sidebar {
        grid-area: sidebar;
    }

    .content {
        grid-area: content;
    }
</style>