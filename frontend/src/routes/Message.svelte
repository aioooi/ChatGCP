<script lang="ts">
	export let msg: t.ReceivedMessage;
</script>

<div class="message-cls">
	{msg.user_id}
	{msg.time}
	{msg.message}
</div>
