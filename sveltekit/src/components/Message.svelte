<script lang="ts">
    interface Message {
        profile: {
            name: string;
            image: string;
        }
        message: string;
        type: 'incoming' | 'outgoing';
    }

    export let message: Message;

    import snarkdown from "snarkdown";
</script>

<div class="message {message.type}">
    <span>
        <img src={message.profile.image} alt={message.profile.name} class="w-8 h-8 rounded-full inline-block mr-2" />
        <span class="font-bold">{message.profile.name}</span>
    </span>

    <span class="ml-12 prose prose-invert text-left">
        {@html snarkdown(message.message)}
    </span>
</div>

<style lang="postcss">

    .message {
        @apply flex items-start flex-col w-full px-5 md:px-[20%] lg:px-[30%] py-[15px];
    }
</style>