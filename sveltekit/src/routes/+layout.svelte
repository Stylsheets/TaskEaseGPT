<script>
    import "../app.css";
    import Sidebar from "../components/Sidebar.svelte";
</script>

<div class="grid-container">
    <aside class="sidebar"><Sidebar/></aside>
    <main class="content">
        <slot/>
    </main>
</div>

<style lang="postcss">

    .grid-container {
        display: grid;
        grid-template-rows: 100%;
        grid-template-columns: 300px 1fr;
        grid-template-areas: "sidebar content";
        height: 100vh;
    }

    .sidebar {
        grid-area: sidebar;
        @apply border-r border-light rounded-r-md;
    }

    .content {
        grid-area: content;
    }
</style>