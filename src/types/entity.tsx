export interface User {
    username: string;
    userId: string;
    isAuthenticated: boolean;
    isLoading: boolean;
};

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
    likes: number;
    dislikes: number;
    user_vote: 1 | -1 | null;
    created_at: string;
    updated_at: string;
}

export interface Comment {
    comment_id: number;
    post_id: number;
    user_id: number;
    username: string;
    description: string;
    likes: number;
    dislikes: number;
    user_vote: 1 | -1 | null;
    created_at: string;
    updated_at: string;
}

export type Entity = Topic | Post | Comment;
