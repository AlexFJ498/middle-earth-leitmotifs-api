CREATE TABLE tracks_themes (
    track_id UUID NOT NULL,
    theme_id UUID NOT NULL,
    start_second INTEGER NOT NULL,
    end_second INTEGER NOT NULL,
    is_variant BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    updated_at TIMESTAMP(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (track_id, theme_id, start_second),
    FOREIGN KEY (track_id) REFERENCES tracks(id) ON DELETE CASCADE,
    FOREIGN KEY (theme_id) REFERENCES themes(id) ON DELETE CASCADE
);