interface OnlineUser {
	id: string;
	name: string;
	color: number;
}

declare enum ReceivedMessageAuthor {
	ThisUser,
	OtherUser,
	System
}

interface ReceivedMessage {
	user_id: string;
	message: string;
	time: string;
	// TODO not really clean to mix in the UI stuff
	color: number;
	user_name?: string;
	authorType: ReceivedMessageAuthor;
}
