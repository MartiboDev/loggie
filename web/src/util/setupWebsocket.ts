export default function setupWebsocket() {
	const ws = new WebSocket(
		`ws://localhost:${import.meta.env.VITE_SERVER_PORT ?? 8080}/ws`
	)

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
