<script lang="ts">
    import type {IConversation} from "../types";
    import Conversation from "./Conversation.svelte";
    import Button from "./Button.svelte";
    import ButtonLink from "./ButtonLink.svelte";

    const conversations: IConversation[] = []

    // Populate conversations for testing
    for (let i = 0; i < 25; i++) {
        conversations.push({
            id: i,
            title: `Conversation ${i}`,
        })
    }
</script>

<div class="sidebar-container">
    <div class="top flex flex-col gap-5 items-center">
        <img alt="Logo" src="/logo.png">
        <button class="flex items-center bg-white pl-8 pr-3 w-[135px] py-1 rounded-full relative">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor"
                 class="w-12 h-12 absolute fill-primary -left-4 ">
                <path fill-rule="evenodd"
                      d="M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25zM12.75 9a.75.75 0 00-1.5 0v2.25H9a.75.75 0 000 1.5h2.25V15a.75.75 0 001.5 0v-2.25H15a.75.75 0 000-1.5h-2.25V9z"
                      clip-rule="evenodd"/>
            </svg>
            <span class="text-base font-semibold text-right text-black">New Chat</span>
        </button>
    </div>
    <ul class="main flex flex-col h-full gap-2 overflow-y-scroll p-2">
        {#each conversations as conversation}
            <li>
                <Conversation {conversation}/>
            </li>
        {/each}
    </ul>
    <ul class="bottom [&>*]:w-full flex gap-2 p-2 flex-col items-center">
        <li>
            <Button>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                    <path fill-rule="evenodd" d="M7.5 6a4.5 4.5 0 119 0 4.5 4.5 0 01-9 0zM3.751 20.105a8.25 8.25 0 0116.498 0 .75.75 0 01-.437.695A18.683 18.683 0 0112 22.5c-2.786 0-5.433-.608-7.812-1.7a.75.75 0 01-.437-.695z" clip-rule="evenodd" />
                </svg>
                <span class="group-hover:scale-105 transition-all">Profile</span>
            </Button>
        </li>

        <li>
            <Button>
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6 group-hover:rotate-180 transition-all">
                    <path fill-rule="evenodd"
                          d="M11.078 2.25c-.917 0-1.699.663-1.85 1.567L9.05 4.889c-.02.12-.115.26-.297.348a7.493 7.493 0 00-.986.57c-.166.115-.334.126-.45.083L6.3 5.508a1.875 1.875 0 00-2.282.819l-.922 1.597a1.875 1.875 0 00.432 2.385l.84.692c.095.078.17.229.154.43a7.598 7.598 0 000 1.139c.015.2-.059.352-.153.43l-.841.692a1.875 1.875 0 00-.432 2.385l.922 1.597a1.875 1.875 0 002.282.818l1.019-.382c.115-.043.283-.031.45.082.312.214.641.405.985.57.182.088.277.228.297.35l.178 1.071c.151.904.933 1.567 1.85 1.567h1.844c.916 0 1.699-.663 1.85-1.567l.178-1.072c.02-.12.114-.26.297-.349.344-.165.673-.356.985-.57.167-.114.335-.125.45-.082l1.02.382a1.875 1.875 0 002.28-.819l.923-1.597a1.875 1.875 0 00-.432-2.385l-.84-.692c-.095-.078-.17-.229-.154-.43a7.614 7.614 0 000-1.139c-.016-.2.059-.352.153-.43l.84-.692c.708-.582.891-1.59.433-2.385l-.922-1.597a1.875 1.875 0 00-2.282-.818l-1.02.382c-.114.043-.282.031-.449-.083a7.49 7.49 0 00-.985-.57c-.183-.087-.277-.227-.297-.348l-.179-1.072a1.875 1.875 0 00-1.85-1.567h-1.843zM12 15.75a3.75 3.75 0 100-7.5 3.75 3.75 0 000 7.5z"
                          clip-rule="evenodd"/>
                </svg>
                <span class="group-hover:scale-105 transition-all">Settings</span>
            </Button>
        </li>

        <li>
            <ButtonLink target="_blank" href="https://github.com/Stylsheets/TaskEaseGPT">
                <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
                    <path fill-rule="evenodd" d="M3 6a3 3 0 013-3h12a3 3 0 013 3v12a3 3 0 01-3 3H6a3 3 0 01-3-3V6zm14.25 6a.75.75 0 01-.22.53l-2.25 2.25a.75.75 0 11-1.06-1.06L15.44 12l-1.72-1.72a.75.75 0 111.06-1.06l2.25 2.25c.141.14.22.331.22.53zm-10.28-.53a.75.75 0 000 1.06l2.25 2.25a.75.75 0 101.06-1.06L8.56 12l1.72-1.72a.75.75 0 10-1.06-1.06l-2.25 2.25z" clip-rule="evenodd" />
                </svg>

                <span class="group-hover:scale-105 transition-all">Code</span>
            </ButtonLink>
        </li>

    </ul>
</div>

<style lang="postcss">
    .sidebar-container {
        display: grid;
        grid-template-rows: 150px 1fr 150px;
        grid-template-columns: 1fr;
        grid-template-areas: "top" "main" "bottom";
        @apply divide-y divide-light h-full;
    }

    .top {
        grid-area: top;
        padding: 12px;
    }

    .main {
        grid-area: main;
    }

    .bottom {
        grid-area: bottom;
    }

    .main::-webkit-scrollbar {
        @apply w-2;
    }

    /* Track */
    .main::-webkit-scrollbar-track {
        @apply bg-accent;
    }

    /* Handle */
    .main::-webkit-scrollbar-thumb {
        @apply bg-primary rounded;
    }

    /* Handle on hover */
    .main::-webkit-scrollbar-thumb:hover {
        @apply bg-secondary;
    }
</style>