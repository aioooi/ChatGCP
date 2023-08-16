declare namespace t {
	interface OnlineUser {
		id: string;
		name: string;
		color: number;
	}

	interface ReceivedMessage {
		user_id?: string;
		author: ReceivedMessageAuthor;
		message: string;
		time: string;
	}

	enum ReceivedMessageAuthor {
		ThisUser,
		OtherUser,
		System
	}
}
