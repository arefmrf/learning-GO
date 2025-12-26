package public_host

import (
	"context"

	phost "trip/internal/persistence/host/query"
	"trip/internal/services/public_host/responses"
)

func GetPublicHosts(ctx context.Context) ([]responses.HostListItem, error) {
	rows, err := phost.FetchPublicHosts(ctx)
	if err != nil {
		return nil, err
	}

	items := make([]responses.HostListItem, 0, len(rows))
	for _, r := range rows {
		items = append(items, responses.HostListItem{
			UID:       r.UID,
			Title:     r.Title,
			Rate:      r.Rate,
			RateCount: r.RateCount,
			Priority:  r.Priority,
		})
	}

	return items, nil
}
