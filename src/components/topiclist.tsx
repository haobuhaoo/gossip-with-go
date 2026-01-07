import React from "react";
import { useNavigate } from "react-router-dom";

import type { Topic } from "../types/entity";

const TopicList: React.FC = () => {
    const topiclist: Topic[] = [
        { id: 1, topicTitle: "Tech" },
        { id: 2, topicTitle: "Food" },
        { id: 3, topicTitle: "Business" },
    ];

    const navigate = useNavigate();

    const handleClick = (t: Topic) => {
        navigate(`/home/${t.id}`, { state: t });
    };

    return (
        <div>
            {topiclist.map((t: Topic) => (
                <div onClick={() => handleClick(t)}>{t.topicTitle}</div>
            ))}
        </div>
    );
};

export default TopicList;
