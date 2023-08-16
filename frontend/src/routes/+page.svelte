<script lang="ts">
	import * as clientState from '$lib/ts/stores';
	import Chat from './Chat.svelte';
	import Login from './Login.svelte';

	let user = '';
	clientState.user.subscribe((v) => {
		user = v;
	});
</script>

{#if user == ''}
	<Login />
{:else}
	<Chat />
{/if}
