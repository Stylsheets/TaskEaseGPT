<script lang="ts">
    export let func: (() => void) | undefined = undefined;
    let clazz: string | undefined = undefined;
    export { clazz as class };
</script>

<button on:click={func} class="{clazz} flex items-center justify-center gap-2 p-2 rounded-lg font-semibold w-full bg-accent hover:mix-blend-plus-lighter active:scale-105 transition-all duration-100 group">
    <slot />
</button>