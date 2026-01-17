export interface Topic {
    topic_id: number;
    title: string;
    user_id: number;
    created_at: string;
}

export interface Post {
    post_id: number;
    topic_id: number;
    user_id: number;
    username: string;
    title: string;
    description: string;
    created_at: string;
    updated_at: string;
}
