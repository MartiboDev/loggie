export default function setupWebsocket() {
	const ws = new WebSocket("ws://localhost:8080/ws")

	ws.onopen = () => {
		console.log("Connected to the server")
	}

	ws.onclose = () => {
		console.log("Disconnected from the server")
	}

	ws.onerror = (error) => {
		console.error("Error:", error)
	}

	return ws
}
