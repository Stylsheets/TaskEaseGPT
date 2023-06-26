<script>
    import "../app.css";
    import Sidebar from "../components/Sidebar.svelte";
    import Navbar from "../components/Navbar.svelte";
</script>

<div class="grid-container">
    <header class="navbar"><Navbar/></header>
    <aside class="sidebar"><Sidebar/></aside>
    <main class="content">
        <slot/>
    </main>
</div>

<style lang="postcss">

    .grid-container {
        display: grid;
        grid-template-columns: 100%;
        grid-template-rows: 60px auto;
        grid-template-areas:
            "navbar"
            "content";
        height: 100vh;
    }

    @media screen and (min-width: 768px) {
        .grid-container {
            grid-template-columns: 250px auto;
            grid-template-rows: 60px auto;
            grid-template-areas:
                 "sidebar navbar"
                 "sidebar content";
        }
    }

    .navbar {
        grid-area: navbar;
        @apply border-b border-border;
    }

    .sidebar {
        grid-area: sidebar;
        @apply border-r border-border;
    }

    .content {
        grid-area: content;
    }
</style>