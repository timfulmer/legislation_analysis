CREATE TABLE legislative_item (
    created TIMESTAMP NOT NULL DEFAULT NOW(),
    updated TIMESTAMP NOT NULL DEFAULT NOW(),
    legislative_text TEXT NOT NULL PRIMARY KEY,
    incidence_count INT NOT NULL,
    found_in_bills TEXT[] NOT NULL DEFAULT '{}'::TEXT[],
    cosponsors TEXT[] NOT NULL DEFAULT '{}'::TEXT[],
    total_incidence_count INT NOT NULL
)