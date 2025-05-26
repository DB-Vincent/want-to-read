export interface Book {
    id?: number;
    title: string;
    author: string;
    user_id?: number;
    created_at?: string;
    updated_at?: string;
    deleted_at?: string | null;  // Optional since it's marked with json:"-" in Go
    completed: boolean;
} 