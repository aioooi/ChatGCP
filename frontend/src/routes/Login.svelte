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

<h1>HELLO</h1>
<h2>MY NAME IS</h2>
<form on:submit|preventDefault={handleSubmit}>
	<label
		>Username:
		<input type="text" name="user" placeholder="<YOUR NAME>" autofocus />
		<!-- <input type="button" value="Login" /> -->
	</label>
</form>

<style>
	h1 {
		color: red;
	}
</style>
