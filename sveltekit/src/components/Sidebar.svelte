<script lang="ts">
    interface Item {
        id: number;
        name: string;
        type: string;
    }

    // Data
    export let conversations: Item[] = [
        {
            id: 1,
            name: 'Conversation 1',
            type: 'conversation',
        },
    ];
    export let knowledgeBases: Item[] = [];

    enum ActivePage {
        Conversations = conversations,
        KnowledgeBases = knowledgeBases,
    }

    let activePage: ActivePage = ActivePage.Conversations;
    let activeItemId: number = conversations[0].id;
</script>

<aside class="w-full flex flex-col divide-y divide-white py-2">
    <nav class="w-full flex flex-col gap-2 items-center pb-2">
        <span class="text-text-primary hover:text-text-secondary cursor-pointer text-xl font-bold">Conversations</span>
        <span class="text-text-primary hover:text-text-secondary cursor-pointer text-xl font-bold">Knowledgebase</span>
    </nav>
    <ul class="flex flex-col w-full text-left pl-2">
        {#each activePage as item}
            <li class="flex items-center gap-2 w-full cursor-pointer text-text-primary hover:text-text-secondary text-lg font-semibold">
                {item.name}
                {#if item.id === activeItemId}
                    <div class="w-2 h-2 rounded-full bg-green-400"></div>
                {/if}
            </li>
        {/each}
    </ul>
</aside>