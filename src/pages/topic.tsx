import React, { useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";

import type { Post } from "../types/entity";
import PostList from "../components/postlist";

const TopicPage: React.FC = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const { id, topicTitle } = location.state || {};

    useEffect(() => {
        if (id) {
            console.log(id);
            console.log(topicTitle);
        }
    }, [id]);

    const handleClick = (p: Post) => {
        navigate(`/home/${id}/${p.id}`, { state: p });
    };

    return (
        <div>
            <h2>Topic: {topicTitle}</h2>
            <PostList handleClick={handleClick} />
        </div>
    );
};

export default TopicPage;
