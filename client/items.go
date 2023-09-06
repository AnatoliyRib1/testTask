package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"golang.org/x/time/rate"
	"testTask/contracts"
)

// GetLimits возвращает количество элементов и временной интервал, ограничивающие обработку внешним сервисом.
func (c *Client) GetLimits() (*contracts.Limits, error) {
	resp, err := c.client.R().Get(c.path("items/limits"))
	if err != nil {
		return nil, err
	}

	// Распаковка ответа в структуру с лимитами
	limits := new(contracts.Limits)
	err = json.Unmarshal(resp.Body(), &limits)
	if err != nil {
		return nil, fmt.Errorf("failed to parse limits response: %s", err)
	}

	return limits, nil
}

// ProcessBatch обрабатывает батч объектов.
func (c *Client) ProcessBatch(ctx context.Context, batch contracts.Batch) error {
	err := c.setLimiter()
	if err != nil {
		return err
	}

	if !c.limiter.Allow() {
		return errors.New("rate limiter exceeded")
	}

	_, err = c.client.R().
		SetContext(ctx).
		SetBody(batch).
		Post(c.path("items/process"))
	return err
}

func (c *Client) setLimiter() error {
	if c.limiter != nil {
		return nil
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// double check
	if c.limiter != nil {
		return nil
	}

	limits, err := c.GetLimits()
	if err != nil {
		return err
	}
	c.limiter = rate.NewLimiter(rate.Every(limits.Interval), limits.MaxItems)
	return nil
}
