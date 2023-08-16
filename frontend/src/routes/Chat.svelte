<script lang="ts">
	import * as clientState from '$lib/ts/stores';
	import Message from './Message.svelte';

	enum MESSAGE_TYPE {
		ACTION_LOGIN = 'action-login',
		EVENT_LOGIN = 'event-login',
		EVENT_USERS_UPDATED = 'event-users-updated',
		PAYLOAD = 'payload'
	}

	// WebSocket connection to backend:
	const BACKEND_HOST = import.meta.env.VITE_CHATGCP_BACKEND_HOST;
	// TODO test for undefined here instead of 'undefined'?
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
				handleUserUpdate(m.payload);
				console.log(ONLINE_USERS);
				break;
			}
			case MESSAGE_TYPE.PAYLOAD: {
				messageBuffer = [
					...messageBuffer,
					{
						...m.payload,
						time: now(),
						author:
							m.payload['user_id'] == USER_ID
								? t.ReceivedMessageAuthor.ThisUser
								: t.ReceivedMessageAuthor.OtherUser
					}
				];
				console.log(messageBuffer);
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
		return String(new Date().toLocaleString().split(', ')[1]);
	}

	// Update ONLINE_USERS
	function handleUserUpdate(updatedUsers: Object) {
		// add new users
		for (const [user_id, user_name] of Object.entries(updatedUsers)) {
			if (!ONLINE_USERS.has(user_id)) {
				ONLINE_USERS.set(user_id, {
					id: user_id,
					name: user_name,
					color: nextColor()
				});
			}
		}
		// remove disconnected users
		for (const [user_id, _] of ONLINE_USERS.entries()) {
			if (!(user_id in updatedUsers)) {
				ONLINE_USERS.delete(user_id);
			}
		}
		// TODO put this in a store?
		ONLINE_USERS = ONLINE_USERS;
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

	let COLORS = 8;
	let color = COLORS - 1;
	const nextColor = () => {
		color += 1;
		return color % COLORS;
	};

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
	let ONLINE_USERS = new Map<string, t.OnlineUser>();

	let messageBuffer: t.ReceivedMessage[] = [];
</script>

<h1>Running in {import.meta.env.VITE_CHATGCP_BACKEND_HOST}</h1>

<div id="container">
	<div id="users-container">
		<div id="users">
			<ul>
				{#each [...ONLINE_USERS.entries()] as [_, u]}
					<li class={`color-${u.color}`}>{u.name}</li>
				{/each}
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
