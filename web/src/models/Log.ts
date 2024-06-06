export interface Log {
	id: number
	source: number
	severity: string
	category: string
	resource: string
	timestamp: Date
	message: string
}
