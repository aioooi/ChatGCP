<script lang="ts">
	import { getStores } from "$app/stores";
  import * as clientState from "$lib/ts/stores";

  const handleSubmit = (e) => {
    let data = new FormData(e.target);
    console.log(data)
    clientState.logged_in.set(true)
    clientState.user.set(data.get("user"))
  }
</script>

<h1>HELLO</h1>
<h2>MY NAME IS</h2>
<form on:submit|preventDefault={handleSubmit}>
  <label>Username:
    <input type="text" name="user" placeholder="<YOUR NAME>" />
    <!-- <input type="button" value="Login" /> -->
  </label>
</form>

<style>
  h1 {
    color: red;
  }
</style>