import { Box, Typography } from "@mui/material";

import type { Comment, Post } from "../types/entity";

import { showLastUpdated, truncate } from "../utils/formatters";

type Props = {
    /**
     * Entity to retrieve information from.
     */
    entity: Comment | Post;
}

/**
 * Renders the author's username and the entity's last updated time.
 */
const DisplayAuthor: React.FC<Props> = ({ entity }) => {
    return (
        <Box sx={{ display: "flex", justifyContent: "center", alignItems: "center" }}>
            <Typography>
                {truncate(entity.username)} • {showLastUpdated(entity.updated_at)}
            </Typography>
        </Box>
    )
}

export default DisplayAuthor;
