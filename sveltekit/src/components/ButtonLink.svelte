<script lang="ts">
    export let href = "";
    export let target = "_self";
    let clazz: string | undefined = undefined;
    export { clazz as class };
</script>

<a target={target} href={href} class="{clazz} flex items-center justify-center gap-2 p-2 rounded-lg font-semibold w-full bg-accent hover:mix-blend-plus-lighter active:scale-105 transition-all duration-100 group">
    <slot />
</a>