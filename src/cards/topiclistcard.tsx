import React from "react";
import { Box, capitalize, Card, CardContent, Typography } from "@mui/material";
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';

import type { Topic } from "../types/entity";
import { truncate } from "../utils/formatters";

type Props = {
    /**
     * Topic to be displayed.
     */
    topic: Topic;

    /**
     * Function that passes selected topic to parent component.
     */
    handleClick: (t: Topic) => void;

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
const TopicListCard: React.FC<Props> = ({ topic, handleClick, openTopicModal, onDelete }) => {
    return (
        <Card
            onClick={() => handleClick(topic)}
            sx={{
                width: "60vw",
                margin: "8px",
                border: "2px solid",
                borderColor: "divider",
                borderRadius: 5,
                boxShadow: "0 1px 3px rgba(0, 0, 0, 0.12), 0 2px 4px rgba(0, 0, 0, 0.08)",
                cursor: "pointer",
                "&:hover": { boxShadow: "0 6px 12px rgba(0, 0, 0, 0.12)" },
                "&:hover div": { opacity: 1 }
            }}>
            <CardContent sx={{ display: "flex", alignItems: "center" }}>
                <Typography sx={{ fontSize: "20px", paddingX: "20px", paddingTop: "8px" }}>
                    {truncate(capitalize(topic.title))}
                </Typography>

                <Box
                    sx={{
                        display: "flex",
                        alignItems: "center",
                        marginLeft: "auto",
                        fontSize: "24px",
                        paddingTop: "8px",
                        paddingRight: "8px",
                        gap: "8px",
                        cursor: "pointer",
                        opacity: 0,
                        transition: "opacity 0.2s ease",
                    }}>
                    <EditIcon
                        onClick={(event: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
                            event.stopPropagation();
                            openTopicModal(topic);
                        }}
                        sx={{ "&:hover": { color: "blue" } }}
                    />
                    <DeleteIcon
                        onClick={(event: React.MouseEvent<SVGSVGElement, MouseEvent>) => {
                            event.stopPropagation();
                            onDelete(topic);
                        }}
                        sx={{ "&:hover": { color: "red" } }}
                    />
                </Box>
            </CardContent>
        </Card>
    )
}

export default TopicListCard;
