<script lang="ts">
	import * as client_state from '$lib/ts/stores';
  import Login from './Login.svelte';

  let logged_in;
  client_state.logged_in.subscribe(v => {
    logged_in = v;
  })

  let user;
  client_state.user.subscribe(v => {
    user = v;
  })

</script>

{#if !logged_in}
  <Login />
{:else}
  <h1>Welcome to SvelteKit{#if logged_in}, {user}{/if}</h1>
  <p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the
  documentation</p>
{/if}
