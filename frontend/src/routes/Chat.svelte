<script lang="ts">
	import * as clientState from '$lib/ts/stores';
	// import * as t from '$lib/ts/types';
	// import type { ChatMessage } from '$lib/ts/types/main';

	import Message from './Message.svelte';

	const MESSAGE_TYPE = {
		ACTION_LOGIN: 'action-login',
		EVENT_LOGIN: 'event-login',
		EVENT_USERS_UPDATED: 'event-users-updated',
		PAYLOAD: 'payload'
	};

	// WebSocket connection to backend:
	const BACKEND_HOST = import.meta.env.VITE_CHATGCP_BACKEND_HOST;
	if (BACKEND_HOST == 'undefined') {
		// TODO redirect to error page
	}

	let ws = new WebSocket(`ws://${BACKEND_HOST}/ws`);

	ws.addEventListener('open', (e) => {
		// TODO: this is brittle: it relies on user being not empty
		connect();
	});

	function sendMessage(msg: string) {
		console.log(`Sending\n${msg}`);
		ws.send(msg);
	}

	// function receive ;)
	ws.onmessage = (msg) => {
		console.log(`Received\n${msg.data}`);
		const m = JSON.parse(msg.data);
		switch (m.type) {
			case MESSAGE_TYPE.EVENT_LOGIN: {
				USER_ID = m.payload['id'];
				console.log(`Logged in, ID=${USER_ID}`);
				break;
			}
			case MESSAGE_TYPE.EVENT_USERS_UPDATED: {
				ONLINE_USERS = m.payload;
				console.log(ONLINE_USERS);
				break;
			}
			case MESSAGE_TYPE.PAYLOAD: {
				messageBuffer = [...messageBuffer, m.payload];
				console.log(messageBuffer);
				// TODO add message payload to chat window
				// var line =
				//   "[" +
				//   now() +
				//   "] " +
				//   ONLINE_USERS[m.payload["user_id"]] +
				//   ": " +
				//   m.payload["message"] +
				//   "\n";
				// chat.innerText += line;
			}
		}
	};

	// Connect client to server
	function connect() {
		console.log('Connecting to server...');
		if (user != '') {
			sendMessage(
				JSON.stringify({
					type: MESSAGE_TYPE.ACTION_LOGIN,
					payload: { user: user }
				})
			);
		}
	}

	// Returns current timestamp as string
	function now() {
		return new Date().toLocaleString().split(', ')[1];
	}

	function handleEnteredMessage(e) {
		let data = new FormData(e.target);
		console.log(data);

		let m = data.get('message');
		if (m != null) {
			m = m.toString();
			if (m != '') {
				sendMessage(
					JSON.stringify({
						type: MESSAGE_TYPE.PAYLOAD,
						payload: { user_id: USER_ID, message: m }
					})
				);
			}
		}
	}

	// Client state
	let user = '';
	clientState.user.subscribe((v) => {
		user = v;
		console.log(user);
		// This is very brittle.
		// TODO restructure code, extract chat protocol implementation in own
		// module
	});

	// TODO
	let USER_ID = '';
	let ONLINE_USERS = {};

	let messageBuffer: any[] = [];
</script>

<h1>Running in {import.meta.env.VITE_CHATGCP_BACKEND_HOST}</h1>

<div id="container">
	<div id="users-container">
		<div id="users">
			<ul>
				<!-- {#each ONLINE_USERS as u}
					<li class={u.color}>{u.name}</li>
				{/each} -->
			</ul>
		</div>
	</div>

	<div id="main-container">
		<div id="message-container-outer">
			<div id="message-container-inner">
				{#each messageBuffer as m}
					<Message msg={m} />
				{/each}
			</div>
			<!-- message-container-inner -->
		</div>
		<!-- message-container-outer -->

		<div id="chat">
			<form on:submit|preventDefault={handleEnteredMessage}>
				<input
					type="text"
					name="message"
					placeholder="Type a message and hit ENTER :)"
				/>
				<!-- TODO button? -->
			</form>
		</div>
		<!-- chat -->
	</div>
	<!-- main-container -->
</div>
<!-- container -->
