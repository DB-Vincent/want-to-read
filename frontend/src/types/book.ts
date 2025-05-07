export interface Book {
    id: number;
    title: string;
    author: string;
    created_at: string;  // ISO date string
    updated_at: string;  // ISO date string
    deleted_at?: string | null;  // Optional since it's marked with json:"-" in Go
    completed: boolean;
} 