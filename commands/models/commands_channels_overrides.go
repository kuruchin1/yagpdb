// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// CommandsChannelsOverride is an object representing the database table.
type CommandsChannelsOverride struct {
	ID                      int64            `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID                 int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Channels                types.Int64Array `boil:"channels" json:"channels,omitempty" toml:"channels" yaml:"channels,omitempty"`
	ChannelCategories       types.Int64Array `boil:"channel_categories" json:"channel_categories,omitempty" toml:"channel_categories" yaml:"channel_categories,omitempty"`
	Global                  bool             `boil:"global" json:"global" toml:"global" yaml:"global"`
	CommandsEnabled         bool             `boil:"commands_enabled" json:"commands_enabled" toml:"commands_enabled" yaml:"commands_enabled"`
	AutodeleteResponse      bool             `boil:"autodelete_response" json:"autodelete_response" toml:"autodelete_response" yaml:"autodelete_response"`
	AutodeleteTrigger       bool             `boil:"autodelete_trigger" json:"autodelete_trigger" toml:"autodelete_trigger" yaml:"autodelete_trigger"`
	AutodeleteResponseDelay int              `boil:"autodelete_response_delay" json:"autodelete_response_delay" toml:"autodelete_response_delay" yaml:"autodelete_response_delay"`
	AutodeleteTriggerDelay  int              `boil:"autodelete_trigger_delay" json:"autodelete_trigger_delay" toml:"autodelete_trigger_delay" yaml:"autodelete_trigger_delay"`
	RequireRoles            types.Int64Array `boil:"require_roles" json:"require_roles" toml:"require_roles" yaml:"require_roles"`
	IgnoreRoles             types.Int64Array `boil:"ignore_roles" json:"ignore_roles" toml:"ignore_roles" yaml:"ignore_roles"`

	R *commandsChannelsOverrideR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L commandsChannelsOverrideL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CommandsChannelsOverrideColumns = struct {
	ID                      string
	GuildID                 string
	Channels                string
	ChannelCategories       string
	Global                  string
	CommandsEnabled         string
	AutodeleteResponse      string
	AutodeleteTrigger       string
	AutodeleteResponseDelay string
	AutodeleteTriggerDelay  string
	RequireRoles            string
	IgnoreRoles             string
}{
	ID:                      "id",
	GuildID:                 "guild_id",
	Channels:                "channels",
	ChannelCategories:       "channel_categories",
	Global:                  "global",
	CommandsEnabled:         "commands_enabled",
	AutodeleteResponse:      "autodelete_response",
	AutodeleteTrigger:       "autodelete_trigger",
	AutodeleteResponseDelay: "autodelete_response_delay",
	AutodeleteTriggerDelay:  "autodelete_trigger_delay",
	RequireRoles:            "require_roles",
	IgnoreRoles:             "ignore_roles",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_Int64Array) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_Int64Array) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var CommandsChannelsOverrideWhere = struct {
	ID                      whereHelperint64
	GuildID                 whereHelperint64
	Channels                whereHelpertypes_Int64Array
	ChannelCategories       whereHelpertypes_Int64Array
	Global                  whereHelperbool
	CommandsEnabled         whereHelperbool
	AutodeleteResponse      whereHelperbool
	AutodeleteTrigger       whereHelperbool
	AutodeleteResponseDelay whereHelperint
	AutodeleteTriggerDelay  whereHelperint
	RequireRoles            whereHelpertypes_Int64Array
	IgnoreRoles             whereHelpertypes_Int64Array
}{
	ID:                      whereHelperint64{field: `id`},
	GuildID:                 whereHelperint64{field: `guild_id`},
	Channels:                whereHelpertypes_Int64Array{field: `channels`},
	ChannelCategories:       whereHelpertypes_Int64Array{field: `channel_categories`},
	Global:                  whereHelperbool{field: `global`},
	CommandsEnabled:         whereHelperbool{field: `commands_enabled`},
	AutodeleteResponse:      whereHelperbool{field: `autodelete_response`},
	AutodeleteTrigger:       whereHelperbool{field: `autodelete_trigger`},
	AutodeleteResponseDelay: whereHelperint{field: `autodelete_response_delay`},
	AutodeleteTriggerDelay:  whereHelperint{field: `autodelete_trigger_delay`},
	RequireRoles:            whereHelpertypes_Int64Array{field: `require_roles`},
	IgnoreRoles:             whereHelpertypes_Int64Array{field: `ignore_roles`},
}

// CommandsChannelsOverrideRels is where relationship names are stored.
var CommandsChannelsOverrideRels = struct {
	CommandsCommandOverrides string
}{
	CommandsCommandOverrides: "CommandsCommandOverrides",
}

// commandsChannelsOverrideR is where relationships are stored.
type commandsChannelsOverrideR struct {
	CommandsCommandOverrides CommandsCommandOverrideSlice
}

// NewStruct creates a new relationship struct
func (*commandsChannelsOverrideR) NewStruct() *commandsChannelsOverrideR {
	return &commandsChannelsOverrideR{}
}

// commandsChannelsOverrideL is where Load methods for each relationship are stored.
type commandsChannelsOverrideL struct{}

var (
	commandsChannelsOverrideColumns               = []string{"id", "guild_id", "channels", "channel_categories", "global", "commands_enabled", "autodelete_response", "autodelete_trigger", "autodelete_response_delay", "autodelete_trigger_delay", "require_roles", "ignore_roles"}
	commandsChannelsOverrideColumnsWithoutDefault = []string{"guild_id", "channels", "channel_categories", "global", "commands_enabled", "autodelete_response", "autodelete_trigger", "autodelete_response_delay", "autodelete_trigger_delay", "require_roles", "ignore_roles"}
	commandsChannelsOverrideColumnsWithDefault    = []string{"id"}
	commandsChannelsOverridePrimaryKeyColumns     = []string{"id"}
)

type (
	// CommandsChannelsOverrideSlice is an alias for a slice of pointers to CommandsChannelsOverride.
	// This should generally be used opposed to []CommandsChannelsOverride.
	CommandsChannelsOverrideSlice []*CommandsChannelsOverride

	commandsChannelsOverrideQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	commandsChannelsOverrideType                 = reflect.TypeOf(&CommandsChannelsOverride{})
	commandsChannelsOverrideMapping              = queries.MakeStructMapping(commandsChannelsOverrideType)
	commandsChannelsOverridePrimaryKeyMapping, _ = queries.BindMapping(commandsChannelsOverrideType, commandsChannelsOverrideMapping, commandsChannelsOverridePrimaryKeyColumns)
	commandsChannelsOverrideInsertCacheMut       sync.RWMutex
	commandsChannelsOverrideInsertCache          = make(map[string]insertCache)
	commandsChannelsOverrideUpdateCacheMut       sync.RWMutex
	commandsChannelsOverrideUpdateCache          = make(map[string]updateCache)
	commandsChannelsOverrideUpsertCacheMut       sync.RWMutex
	commandsChannelsOverrideUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single commandsChannelsOverride record from the query using the global executor.
func (q commandsChannelsOverrideQuery) OneG(ctx context.Context) (*CommandsChannelsOverride, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single commandsChannelsOverride record from the query.
func (q commandsChannelsOverrideQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CommandsChannelsOverride, error) {
	o := &CommandsChannelsOverride{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for commands_channels_overrides")
	}

	return o, nil
}

// AllG returns all CommandsChannelsOverride records from the query using the global executor.
func (q commandsChannelsOverrideQuery) AllG(ctx context.Context) (CommandsChannelsOverrideSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all CommandsChannelsOverride records from the query.
func (q commandsChannelsOverrideQuery) All(ctx context.Context, exec boil.ContextExecutor) (CommandsChannelsOverrideSlice, error) {
	var o []*CommandsChannelsOverride

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CommandsChannelsOverride slice")
	}

	return o, nil
}

// CountG returns the count of all CommandsChannelsOverride records in the query, and panics on error.
func (q commandsChannelsOverrideQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all CommandsChannelsOverride records in the query.
func (q commandsChannelsOverrideQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count commands_channels_overrides rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q commandsChannelsOverrideQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q commandsChannelsOverrideQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if commands_channels_overrides exists")
	}

	return count > 0, nil
}

// CommandsCommandOverrides retrieves all the commands_command_override's CommandsCommandOverrides with an executor.
func (o *CommandsChannelsOverride) CommandsCommandOverrides(mods ...qm.QueryMod) commandsCommandOverrideQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"commands_command_overrides\".\"commands_channels_overrides_id\"=?", o.ID),
	)

	query := CommandsCommandOverrides(queryMods...)
	queries.SetFrom(query.Query, "\"commands_command_overrides\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"commands_command_overrides\".*"})
	}

	return query
}

// LoadCommandsCommandOverrides allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (commandsChannelsOverrideL) LoadCommandsCommandOverrides(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCommandsChannelsOverride interface{}, mods queries.Applicator) error {
	var slice []*CommandsChannelsOverride
	var object *CommandsChannelsOverride

	if singular {
		object = maybeCommandsChannelsOverride.(*CommandsChannelsOverride)
	} else {
		slice = *maybeCommandsChannelsOverride.(*[]*CommandsChannelsOverride)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &commandsChannelsOverrideR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &commandsChannelsOverrideR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`commands_command_overrides`), qm.WhereIn(`commands_channels_overrides_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load commands_command_overrides")
	}

	var resultSlice []*CommandsCommandOverride
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice commands_command_overrides")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on commands_command_overrides")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for commands_command_overrides")
	}

	if singular {
		object.R.CommandsCommandOverrides = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &commandsCommandOverrideR{}
			}
			foreign.R.CommandsChannelsOverride = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.CommandsChannelsOverridesID {
				local.R.CommandsCommandOverrides = append(local.R.CommandsCommandOverrides, foreign)
				if foreign.R == nil {
					foreign.R = &commandsCommandOverrideR{}
				}
				foreign.R.CommandsChannelsOverride = local
				break
			}
		}
	}

	return nil
}

// AddCommandsCommandOverridesG adds the given related objects to the existing relationships
// of the commands_channels_override, optionally inserting them as new records.
// Appends related to o.R.CommandsCommandOverrides.
// Sets related.R.CommandsChannelsOverride appropriately.
// Uses the global database handle.
func (o *CommandsChannelsOverride) AddCommandsCommandOverridesG(ctx context.Context, insert bool, related ...*CommandsCommandOverride) error {
	return o.AddCommandsCommandOverrides(ctx, boil.GetContextDB(), insert, related...)
}

// AddCommandsCommandOverrides adds the given related objects to the existing relationships
// of the commands_channels_override, optionally inserting them as new records.
// Appends related to o.R.CommandsCommandOverrides.
// Sets related.R.CommandsChannelsOverride appropriately.
func (o *CommandsChannelsOverride) AddCommandsCommandOverrides(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*CommandsCommandOverride) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CommandsChannelsOverridesID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"commands_command_overrides\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"commands_channels_overrides_id"}),
				strmangle.WhereClause("\"", "\"", 2, commandsCommandOverridePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CommandsChannelsOverridesID = o.ID
		}
	}

	if o.R == nil {
		o.R = &commandsChannelsOverrideR{
			CommandsCommandOverrides: related,
		}
	} else {
		o.R.CommandsCommandOverrides = append(o.R.CommandsCommandOverrides, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &commandsCommandOverrideR{
				CommandsChannelsOverride: o,
			}
		} else {
			rel.R.CommandsChannelsOverride = o
		}
	}
	return nil
}

// CommandsChannelsOverrides retrieves all the records using an executor.
func CommandsChannelsOverrides(mods ...qm.QueryMod) commandsChannelsOverrideQuery {
	mods = append(mods, qm.From("\"commands_channels_overrides\""))
	return commandsChannelsOverrideQuery{NewQuery(mods...)}
}

// FindCommandsChannelsOverrideG retrieves a single record by ID.
func FindCommandsChannelsOverrideG(ctx context.Context, iD int64, selectCols ...string) (*CommandsChannelsOverride, error) {
	return FindCommandsChannelsOverride(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindCommandsChannelsOverride retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCommandsChannelsOverride(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*CommandsChannelsOverride, error) {
	commandsChannelsOverrideObj := &CommandsChannelsOverride{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"commands_channels_overrides\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, commandsChannelsOverrideObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from commands_channels_overrides")
	}

	return commandsChannelsOverrideObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CommandsChannelsOverride) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CommandsChannelsOverride) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commands_channels_overrides provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(commandsChannelsOverrideColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	commandsChannelsOverrideInsertCacheMut.RLock()
	cache, cached := commandsChannelsOverrideInsertCache[key]
	commandsChannelsOverrideInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			commandsChannelsOverrideColumns,
			commandsChannelsOverrideColumnsWithDefault,
			commandsChannelsOverrideColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(commandsChannelsOverrideType, commandsChannelsOverrideMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(commandsChannelsOverrideType, commandsChannelsOverrideMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"commands_channels_overrides\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"commands_channels_overrides\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into commands_channels_overrides")
	}

	if !cached {
		commandsChannelsOverrideInsertCacheMut.Lock()
		commandsChannelsOverrideInsertCache[key] = cache
		commandsChannelsOverrideInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single CommandsChannelsOverride record using the global executor.
// See Update for more documentation.
func (o *CommandsChannelsOverride) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the CommandsChannelsOverride.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CommandsChannelsOverride) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	commandsChannelsOverrideUpdateCacheMut.RLock()
	cache, cached := commandsChannelsOverrideUpdateCache[key]
	commandsChannelsOverrideUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			commandsChannelsOverrideColumns,
			commandsChannelsOverridePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update commands_channels_overrides, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"commands_channels_overrides\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, commandsChannelsOverridePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(commandsChannelsOverrideType, commandsChannelsOverrideMapping, append(wl, commandsChannelsOverridePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update commands_channels_overrides row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for commands_channels_overrides")
	}

	if !cached {
		commandsChannelsOverrideUpdateCacheMut.Lock()
		commandsChannelsOverrideUpdateCache[key] = cache
		commandsChannelsOverrideUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q commandsChannelsOverrideQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q commandsChannelsOverrideQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for commands_channels_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for commands_channels_overrides")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CommandsChannelsOverrideSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CommandsChannelsOverrideSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commandsChannelsOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"commands_channels_overrides\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, commandsChannelsOverridePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in commandsChannelsOverride slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all commandsChannelsOverride")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CommandsChannelsOverride) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CommandsChannelsOverride) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no commands_channels_overrides provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(commandsChannelsOverrideColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	commandsChannelsOverrideUpsertCacheMut.RLock()
	cache, cached := commandsChannelsOverrideUpsertCache[key]
	commandsChannelsOverrideUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			commandsChannelsOverrideColumns,
			commandsChannelsOverrideColumnsWithDefault,
			commandsChannelsOverrideColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			commandsChannelsOverrideColumns,
			commandsChannelsOverridePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert commands_channels_overrides, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(commandsChannelsOverridePrimaryKeyColumns))
			copy(conflict, commandsChannelsOverridePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"commands_channels_overrides\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(commandsChannelsOverrideType, commandsChannelsOverrideMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(commandsChannelsOverrideType, commandsChannelsOverrideMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert commands_channels_overrides")
	}

	if !cached {
		commandsChannelsOverrideUpsertCacheMut.Lock()
		commandsChannelsOverrideUpsertCache[key] = cache
		commandsChannelsOverrideUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single CommandsChannelsOverride record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CommandsChannelsOverride) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single CommandsChannelsOverride record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CommandsChannelsOverride) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CommandsChannelsOverride provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), commandsChannelsOverridePrimaryKeyMapping)
	sql := "DELETE FROM \"commands_channels_overrides\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from commands_channels_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for commands_channels_overrides")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q commandsChannelsOverrideQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no commandsChannelsOverrideQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commands_channels_overrides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commands_channels_overrides")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o CommandsChannelsOverrideSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CommandsChannelsOverrideSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CommandsChannelsOverride slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commandsChannelsOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"commands_channels_overrides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commandsChannelsOverridePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from commandsChannelsOverride slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for commands_channels_overrides")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CommandsChannelsOverride) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no CommandsChannelsOverride provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CommandsChannelsOverride) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCommandsChannelsOverride(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommandsChannelsOverrideSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty CommandsChannelsOverrideSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CommandsChannelsOverrideSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CommandsChannelsOverrideSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), commandsChannelsOverridePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"commands_channels_overrides\".* FROM \"commands_channels_overrides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, commandsChannelsOverridePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CommandsChannelsOverrideSlice")
	}

	*o = slice

	return nil
}

// CommandsChannelsOverrideExistsG checks if the CommandsChannelsOverride row exists.
func CommandsChannelsOverrideExistsG(ctx context.Context, iD int64) (bool, error) {
	return CommandsChannelsOverrideExists(ctx, boil.GetContextDB(), iD)
}

// CommandsChannelsOverrideExists checks if the CommandsChannelsOverride row exists.
func CommandsChannelsOverrideExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"commands_channels_overrides\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if commands_channels_overrides exists")
	}

	return exists, nil
}
