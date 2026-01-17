import React from "react";
import { Box, capitalize, Card, CardContent, Typography } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';

import type { Topic } from "../types/entity";

type Props = {
    /**
     * Topic to be displayed.
     */
    topic: Topic;

    /**
     * True if the user is the author of the topic.
     */
    isUser: boolean;

    /**
     * Function that passes selected topic id to parent component.
     */
    handleClick: (id: number) => void;

    /**
     * Function that passes topic to be updated to parent component.
     */
    openTopicModal: (t: Topic) => void;

    /**
     * Function that passes topic to be deleted to parent component.
     */
    onDelete: (t: Topic) => void;
}

/**
 * Renders a single topic inside a card, which calls `handleClick`, `openTopicModal` and
 * `onDelete` when clicked.
 */
const TopicListCard: React.FC<Props> = ({ topic, isUser, handleClick, openTopicModal, onDelete }) => {
    return (
        <Card
            onClick={() => handleClick(topic.topic_id)}
            sx={{
                width: "60vw",
                margin: "8px",
                border: "2px solid",
                borderColor: "divider",
                borderRadius: 5,
                boxShadow: "0 1px 3px rgba(0, 0, 0, 0.12), 0 2px 4px rgba(0, 0, 0, 0.08)",
                cursor: "pointer",
                "&:hover": { boxShadow: "0 6px 12px rgba(0, 0, 0, 0.12)" },
            }}>
            <CardContent sx={{ display: "flex", alignItems: "center" }}>
                <Typography sx={{
                    fontSize: "20px",
                    paddingX: "20px",
                    paddingTop: "8px",
                    textOverflow: "ellipsis",
                    overflow: "hidden",
                    display: '-webkit-box',
                    WebkitLineClamp: 1,
                    WebkitBoxOrient: 'vertical',
                }}>
                    {capitalize(topic.title)}
                </Typography>

                {isUser &&
                    <Box
                        sx={{
                            display: "flex",
                            alignItems: "center",
                            marginLeft: "auto",
                            fontSize: "24px",
                            paddingRight: "8px",
                            gap: "4px",
                        }}>
                        <EditIcon
                            onClick={(event: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
                                event.stopPropagation();
                                openTopicModal(topic);
                            }}
                            sx={{
                                p: "3px",
                                borderRadius: 10,
                                color: "blue",
                                "&:hover": { backgroundColor: "lightgrey" }
                            }}
                        />
                        <DeleteIcon
                            onClick={(event: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
                                event.stopPropagation();
                                onDelete(topic);
                            }}
                            sx={{
                                p: "3px",
                                borderRadius: 10,
                                color: "red",
                                "&:hover": { backgroundColor: "lightgrey" }
                            }}
                        />
                    </Box>}
            </CardContent>
        </Card>
    )
}

export default TopicListCard;
