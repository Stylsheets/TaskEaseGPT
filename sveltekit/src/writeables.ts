import { writable } from "svelte/store";
import {browser} from "$app/environment";

export const isSidebarOpen = writable(browser ? window.innerWidth > 1024 : false);