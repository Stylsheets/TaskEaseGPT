<script lang="ts">
    interface Message {
        profile: {
            name: string;
            image: string;
        }
        message: string;
        type: 'incoming' | 'outgoing';
    }

    export let message;
</script>

<div class="message {message.type} {message.type === 'outgoing' ? 'mr-5' : 'ml-5'}">
    <span>
        <img src={message.profile.image} alt={message.profile.name} class="w-8 h-8 rounded-full inline-block mr-2" />
        <span class="font-bold">{message.profile.name}</span>
    </span>

    <span class="triangle {message.type}">
        {message.message}
    </span>
</div>

<style lang="postcss">

    .message {
        @apply flex flex-col p-5 gap-3 bg-selection rounded-xl text-white text-base w-fit max-w-[50%] relative;
    }

    .incoming {
        @apply self-start;
    }

    .outgoing {
        @apply self-end;
    }

    .triangle:after {
        content: ' ';
        position: absolute;
        width: 0;
        height: 0;
        top: 0px;
        bottom: auto;
        border: 20px solid;
        border-color: theme(colors.selection) transparent transparent transparent;
    }

    .incoming:after {
        left: -20px;
        right: auto;
    }

    .outgoing:after {
        left: auto;
        right: -20px;
    }

</style>