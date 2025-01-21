CREATE TABLE block_explorer."LastRefreshViews"
(
    last_refresh timestamp with time zone DEFAULT NOW() NOT NULL
);

CREATE OR REPLACE FUNCTION block_explorer.UpdateLastRefresh()
    RETURNS void AS
$$
BEGIN
    LOCK TABLE block_explorer."LastRefreshViews" IN EXCLUSIVE MODE;

    -- If theres a row, update it
    UPDATE block_explorer."LastRefreshViews" SET last_refresh = NOW();

    -- If no row was updated, insert a new one
    IF NOT FOUND THEN
        INSERT INTO block_explorer."LastRefreshViews" (last_refresh) VALUES (NOW());
    END IF;
END;
$$ LANGUAGE plpgsql;
