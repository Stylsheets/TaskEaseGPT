<script lang="ts">
    import Button from "../components/Button.svelte";

    const isDefault = true;

    let text: string = "";
    $: rowCount =  [...text.matchAll(/\n/g)].length + 1 > 5 ? 5 : [...text.matchAll(/\n/g)].length + 1;

    function sendMessage() {
        // TODO: Send message to server
    }
    $: isValid = text.length > 0;
</script>

<main class="container h-full flex flex-col mx-auto relative">
    {#if isDefault}
        <div class="flex flex-col gap-10 justify-center items-center h-full">
            <img alt="Logo" src="logo.png">
            <span class="max-w-[500px] font-semibold text-center">Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ab asperiores, aut delectus deserunt, dolore esse ex expedita facilis fugit minima nam non numquam quae saepe voluptates! Delectus odit sunt tempore?</span>
            <div class="flex flex-wrap gap-4">
                <Button class="text-sm px-4">
                    Give me ideas
                </Button>
                <Button class="text-sm px-4">
                    Write an email
                </Button>
                <Button class="text-sm px-4">
                    Brainstorm ideas
                </Button>
            </div>
        </div>
    {/if}

    <footer>
        <textarea bind:value={text} rows={rowCount} placeholder="Write your prompt.."
                  class="px-2 font-medium py-2 w-full bg-transparent resize-none outline-0 overflow-y-scroll h-max"
        />
        <button on:click={sendMessage} class="group right-2 absolute">
            <svg class="w-8 p-1 transition-colors rounded-lg {isValid ? 'fill-white bg-green-500' : 'fill-border'}" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
                <path d="M1.94631 9.31555C1.42377 9.14137 1.41965 8.86034 1.95706 8.6812L21.0433 2.31913C21.5717 2.14297 21.8748 2.43878 21.7268 2.95706L16.2736 22.0433C16.1226 22.5718 15.8179 22.5901 15.5946 22.0877L12.0002 14.0002L18.0002 6.00017L10.0002 12.0002L1.94631 9.31555Z"></path>
            </svg>
        </button>
    </footer>
</main>

<style lang="postcss">


    footer {
        @apply absolute flex items-center left-0 right-0 max-w-[600px] bottom-2 bg-accent mx-auto rounded-lg pr-8 border transition-colors border-transparent;
    }

    footer:focus-within {
        @apply border-primary;
    }

    textarea::-webkit-scrollbar {
        width: 5px;
    }

    textarea::-webkit-scrollbar-track {
        background-color: transparent;
    }

    textarea::-webkit-scrollbar-thumb {
        background-color: theme('colors.gray.400');
        border-radius: 5px;
    }

    textarea::-webkit-scrollbar-thumb:hover {
        background-color: theme('colors.gray.600');
        transition: 0.5s;
    }
</style>