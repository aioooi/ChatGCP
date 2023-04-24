import { writable } from "svelte/store";

export const logged_in = writable(false);
export const user = writable("");