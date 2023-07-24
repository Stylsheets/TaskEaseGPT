<script lang="ts">
    import Message from "../components/Message.svelte";

    let messages = [
        {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: 'a',
            type: 'incoming'
        },
        {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: 'a',
            type: 'outgoing'
        },
        {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus adipisci aperiam at dolor dolores, eius eos, ex excepturi in libero mollitia necessitatibus nulla omnis provident quaerat qui quos tempore voluptate.',
            type: 'incoming'
        },
        {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus adipisci aperiam at dolor dolores, eius eos, ex excepturi in libero mollitia necessitatibus nulla omnis provident quaerat qui quos tempore voluptate.',
            type: 'outgoing'
        },
        {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus adipisci aperiam at dolor dolores, eius eos, ex excepturi in libero mollitia necessitatibus nulla omnis provident quaerat qui quos tempore voluptate.',
            type: 'incoming'
        },
        {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Accusamus adipisci aperiam at dolor dolores, eius eos, ex excepturi in libero mollitia necessitatibus nulla omnis provident quaerat qui quos tempore voluptate.',
            type: 'outgoing'
        }

    ]

    // TODO: Scrolls to top of the last child element
    function scrollToBottom() {
        const main = document.getElementById('main');
        main.scrollTop = main.scrollHeight;
    }

    function addMessage(message) {
        messages = [...messages, message];
        scrollToBottom();
    }

    function sendMessage() {
        if (isValid) {
            const message = {
                profile: {
                    name: 'Chat GPT',
                    image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
                },
                message: textarea,
                type: 'outgoing'
            };
            addMessage(message);
            textarea = "";
        } else {
            const footer = document.querySelector('footer');
            footer.classList.add('shake');
            setTimeout(() => {
                footer.classList.remove('shake');
            }, 500);
        }
    }

    // TODO: Send message on enter

    let textarea = "";
    $: rowCount = [...textarea.matchAll(/\n/g)].length + 1 > 5 ? 5 : [...textarea.matchAll(/\n/g)].length + 1;
    $: isValid = textarea.length > 0;
</script>

<section id="main" class="flex flex-col overflow-y-scroll gap-5 w-full h-full pb-[8%] divide-y">
    {#each messages as message}
        <Message {message} />
    {/each}
</section>

<footer>
    <textarea bind:value={textarea} rows={rowCount}></textarea>
    <button on:click={sendMessage} class="group right-2 absolute">
        <svg class="w-8 p-1 rounded-lg {isValid ? 'fill-white bg-green-500' : 'fill-border'}" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
            <path d="M1.94631 9.31555C1.42377 9.14137 1.41965 8.86034 1.95706 8.6812L21.0433 2.31913C21.5717 2.14297 21.8748 2.43878 21.7268 2.95706L16.2736 22.0433C16.1226 22.5718 15.8179 22.5901 15.5946 22.0877L12.0002 14.0002L18.0002 6.00017L10.0002 12.0002L1.94631 9.31555Z"></path>
        </svg>
    </button>

</footer>

<style lang="postcss">
    footer {
        @apply absolute flex items-center left-[250px] right-0 w-[800px] bottom-2 bg-selection mx-auto rounded-lg pr-8;
    }

    textarea {
        @apply text-white px-2 text-lg py-2 w-full bg-transparent resize-none outline-0 focus:border-border overflow-y-scroll h-max;
    }

    /* Set height screen height minus 60px */
    #main {
        @apply h-[calc(100vh-60px)] scroll-smooth;
    }

    /* Scrollbar styling for main content*/
    #main::-webkit-scrollbar {
        width: 5px;
    }

    #main::-webkit-scrollbar-track {
        background-color: transparent;
    }

    #main::-webkit-scrollbar-thumb {
        background-color: theme('colors.gray.400');
        border-radius: 5px;
    }

    #main::-webkit-scrollbar-thumb:hover {
        background-color: theme('colors.gray.600');
        transition: 0.5s;
    }

    /* Scrollbar styling for textarea*/
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
