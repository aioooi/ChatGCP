<script lang="ts">
	import * as clientState from '$lib/ts/stores';

	const handleSubmit = (e) => {
		let data = new FormData(e.target);
		console.log(data);

		let user = data.get('user');
		if (user != null) {
			user = user.toString();
			if (user != '') {
				clientState.user.set(user);
			}
		}
	};
</script>

<!--
  SvelteKit doesn't render the body style from .default./style.css as
  included in ./+layout.svelte without this one:
-->
<div style="height:1px" />

<div id="sticker">
	<h1>HELLO</h1>
	<h2>MY NAME IS</h2>
	<div id="field">
		<form on:submit|preventDefault={handleSubmit} autocomplete="off">
			<input type="text" name="user" placeholder="<YOUR NAME>" autofocus />
		</form>
	</div>
</div>

<style>
	#sticker {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%) translateZ(1px) rotate(-3deg);
		border-radius: 1.5em;
		border: 3px solid red;
		background-color: red;
		padding: 1.5em 0.1em;
		filter: drop-shadow(-3px 3px 2px #000000ff);
	}

	#field {
		background-color: white;
	}

	input[type='text'] {
		position: inherit;
		width: 20ex;
		padding: 0;
		margin: 0;
		border: none;
		padding: 0.2em;
		text-align: center;
		height: 3.5em;
		font-size: 4em;
		color: black;
		background: none;
		transform: rotate(1.1deg);
	}

	input[type='text']:focus {
		outline: none;
	}

	h1,
	h2 {
		text-align: center;
		background-color: red;
		color: white;
		margin: 0;
		padding: 0 0 0.3em;
	}

	h1 {
		font-size: 3em;
	}

	h2 {
		font-size: 1.6em;
	}
</style>
