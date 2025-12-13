import React, { useEffect } from "react";
import { useLocation } from "react-router-dom";

const PostPage: React.FC = () => {
    const location = useLocation();
    const { id, postTitle, postDesc } = location.state || {};

    useEffect(() => {
        if (id) {
            console.log(id);
            console.log(postTitle);
            console.log(postDesc);
        }
    }, [id]);

    return (
        <div>
            <h2>Post Title: {postTitle}</h2>
            <h4>{postDesc}</h4>
        </div>
    );
};

export default PostPage;
