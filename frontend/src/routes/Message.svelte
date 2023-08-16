<script lang="ts">
	import { ReceivedMessageAuthor } from '../types';

	export let msg: ReceivedMessage;
</script>

<div
	class="message"
	class:own={msg.authorType == ReceivedMessageAuthor.ThisUser}
	class:other={msg.authorType == ReceivedMessageAuthor.OtherUser}
>
	{#if msg.authorType == ReceivedMessageAuthor.OtherUser}
		<div class="author {`color-${msg.color}`}">{msg.user_name}</div>
	{/if}
	<div class="message-wrapper">
		<!--
    Extra enclosing div as in
    https://css-tricks.com/float-an-element-to-the-bottom-corner/#aa-markup-and-layout
  -->
		<div>
			<div class="timestamp">{msg.time}</div>
			<p>{msg.message}</p>
		</div>
	</div>
</div>
