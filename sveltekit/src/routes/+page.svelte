<script>
    import Message from "../components/Message.svelte";

    let messages = [
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

    function scrollToBottom() {
        const main = document.querySelector('main');
        main.scrollTop = main.scrollHeight;
    }

    function addMessage(message) {
        messages = [...messages, message];
        scrollToBottom();
    }

    function sendMessage() {
        const input = document.querySelector('input');

        if (input.value === '') {
            const footer = document.querySelector('footer');
            footer.classList.add('shake');
            setTimeout(() => {
                footer.classList.remove('shake');
            }, 300)
            return;
        }

        const message = {
            profile: {
                name: 'Chat GPT',
                image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/0/04/ChatGPT_logo.svg/1200px-ChatGPT_logo.svg.png'
            },
            message: input.value,
            type: 'outgoing'
        }
        addMessage(message);
        input.value = '';
    }

    function checkEnter(event) {
        if (event.key === 'Enter') {
            sendMessage();
        }
    }
</script>

<main class="flex flex-col overflow-y-scroll gap-5 w-full h-full pb-[8%] pt-[2%] px-5">
    {#each messages as message}
        <Message {message} />
    {/each}
</main>

<footer>
    <input on:keydown={checkEnter} type="text" placeholder="Type your prompt" class="w-full md:w-1/3 min-h-fit py-3 pl-4 rounded-l-full bg-selection transition-colors duration-300 peer outline-0 border-r-0 border border-transparent focus:border-white">
    <button on:click={sendMessage} class="relative flex items-center bg-selection rounded-r-full pr-5 hover:text-green-400 transition-colors outline-0 duration-300 border-l-0 border border-transparent peer-focus:border-white">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
             class="w-7 h-7">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M12.75 15l3-3m0 0l-3-3m3 3h-7.5M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
    </button>
</footer>

<style lang="postcss">

    footer {
        @apply absolute left-0 right-0 bottom-5 flex justify-center;
    }

    /* push footer to right */
    @media screen and (min-width: 768px) {
        footer {
            margin-left: 250px;
        }
    }

    /* Set height screen height minus 60px */
    main {
        @apply h-[calc(100vh-60px)];
    }

    /* Scrollbar styling */
    main::-webkit-scrollbar {
        width: 5px;
    }

    main::-webkit-scrollbar-track {
        background-color: transparent;
    }

    main::-webkit-scrollbar-thumb {
        background-color: theme('colors.gray.400');
        border-radius: 5px;
    }

    main::-webkit-scrollbar-thumb:hover {
        background-color: theme('colors.gray.600');
        transition: 0.5s;
    }
</style>