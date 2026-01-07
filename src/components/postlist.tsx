import React from "react";

import type { Post } from "../types/entity";

type Props = {
    handleClick: (p: Post) => void;
};

const PostList: React.FC<Props> = (props) => {
    const postlist: Post[] = [
        { id: 1, postTitle: "hihi", postDesc: "hiihihihihiihih" },
        { id: 2, postTitle: "ahdhs", postDesc: "foooodosodoosodos" },
        { id: 3, postTitle: "obsdb", postDesc: "sdsdsdsdsds" },
    ];

    return (
        <div>
            {postlist.map((p: Post) => (
                <div key={p.id} onClick={() => props.handleClick(p)}>
                    {p.postTitle}
                </div>
            ))}
        </div>
    );
};

export default PostList;
