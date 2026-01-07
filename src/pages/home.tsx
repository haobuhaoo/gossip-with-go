import React from "react";

import TopicList from "../components/topiclist";

const HomePage: React.FC = () => {
    return (
        <div>
            <h2>Choose from 1 of the topics below to share your insights.</h2>
            <TopicList />
        </div>
    );
};

export default HomePage;
