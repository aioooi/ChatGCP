<script lang="ts">
	import { ReceivedMessageAuthor } from '../types';

	export let msg: ReceivedMessage;

	// Extracts HH:MM from timestamp
	function fmt_timestamp(timestamp: Date) {
		//                            "HH:MM:SS".slice(0, 5)
		return timestamp.toLocaleString().split(', ')[1].slice(0, 5);
	}
</script>

<div
	class="message"
	class:own={msg.authorType == ReceivedMessageAuthor.ThisUser}
	class:other={msg.authorType == ReceivedMessageAuthor.OtherUser}
	class:system={msg.authorType == ReceivedMessageAuthor.System}
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
			<span class="timestamp">{fmt_timestamp(msg.timestamp)}</span>
			{#if msg.authorType == ReceivedMessageAuthor.System}
				<span class="author {`user-color-${msg.color}`}">{msg.user_name}</span>
			{/if}
			{msg.message}
		</div>
	</div>
</div>

<style>
	.message {
		max-width: 60%;
		width: fit-content;
		background-color: #4e4e4e;
		padding: 0.6em;
		margin-top: 0.6em;
		margin-bottom: 0.1em;
		color: white;
		border-radius: 0.5em;
		filter: drop-shadow(-3px 3px 2px #00000019);
	}

	/*
  Make timestamp float to right bottom corner::after
  https://css-tricks.com/float-an-element-to-the-bottom-corner/#aa-markup-and-layout
  */
	.message-wrapper {
		display: flex;
		font-size: 1em;
	}

	.timestamp {
		float: right;
		height: 100%;
		display: flex;
		align-items: flex-end;
		/* 100% [height] - [.message-wrapper.font-size as offset] */
		shape-outside: inset(calc(100% - 1em) 0 0);
		color: #989898;
		font-size: 0.7em;
		margin-right: 0.1em;
		margin-left: 0.6em;
	}

	.own {
		margin-left: auto;
		margin-right: 0;
	}

	.other {
		margin-left: 0;
		margin-right: auto;
	}

	.system {
		margin-left: auto;
		margin-right: auto;
	}

	.author {
		font-size: 1em;
		font-weight: 600;
		margin: 0.1em 0 0.2em;
	}
</style>
