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

// SoundboardSound is an object representing the database table.
type SoundboardSound struct {
	ID               int              `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt        time.Time        `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time        `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	GuildID          int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	RequiredRole     string           `boil:"required_role" json:"required_role" toml:"required_role" yaml:"required_role"`
	Name             string           `boil:"name" json:"name" toml:"name" yaml:"name"`
	Status           int              `boil:"status" json:"status" toml:"status" yaml:"status"`
	RequiredRoles    types.Int64Array `boil:"required_roles" json:"required_roles,omitempty" toml:"required_roles" yaml:"required_roles,omitempty"`
	BlacklistedRoles types.Int64Array `boil:"blacklisted_roles" json:"blacklisted_roles,omitempty" toml:"blacklisted_roles" yaml:"blacklisted_roles,omitempty"`

	R *soundboardSoundR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L soundboardSoundL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SoundboardSoundColumns = struct {
	ID               string
	CreatedAt        string
	UpdatedAt        string
	GuildID          string
	RequiredRole     string
	Name             string
	Status           string
	RequiredRoles    string
	BlacklistedRoles string
}{
	ID:               "id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	GuildID:          "guild_id",
	RequiredRole:     "required_role",
	Name:             "name",
	Status:           "status",
	RequiredRoles:    "required_roles",
	BlacklistedRoles: "blacklisted_roles",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

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

var SoundboardSoundWhere = struct {
	ID               whereHelperint
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
	GuildID          whereHelperint64
	RequiredRole     whereHelperstring
	Name             whereHelperstring
	Status           whereHelperint
	RequiredRoles    whereHelpertypes_Int64Array
	BlacklistedRoles whereHelpertypes_Int64Array
}{
	ID:               whereHelperint{field: `id`},
	CreatedAt:        whereHelpertime_Time{field: `created_at`},
	UpdatedAt:        whereHelpertime_Time{field: `updated_at`},
	GuildID:          whereHelperint64{field: `guild_id`},
	RequiredRole:     whereHelperstring{field: `required_role`},
	Name:             whereHelperstring{field: `name`},
	Status:           whereHelperint{field: `status`},
	RequiredRoles:    whereHelpertypes_Int64Array{field: `required_roles`},
	BlacklistedRoles: whereHelpertypes_Int64Array{field: `blacklisted_roles`},
}

// SoundboardSoundRels is where relationship names are stored.
var SoundboardSoundRels = struct {
}{}

// soundboardSoundR is where relationships are stored.
type soundboardSoundR struct {
}

// NewStruct creates a new relationship struct
func (*soundboardSoundR) NewStruct() *soundboardSoundR {
	return &soundboardSoundR{}
}

// soundboardSoundL is where Load methods for each relationship are stored.
type soundboardSoundL struct{}

var (
	soundboardSoundColumns               = []string{"id", "created_at", "updated_at", "guild_id", "required_role", "name", "status", "required_roles", "blacklisted_roles"}
	soundboardSoundColumnsWithoutDefault = []string{"created_at", "updated_at", "guild_id", "required_role", "name", "status", "required_roles", "blacklisted_roles"}
	soundboardSoundColumnsWithDefault    = []string{"id"}
	soundboardSoundPrimaryKeyColumns     = []string{"id"}
)

type (
	// SoundboardSoundSlice is an alias for a slice of pointers to SoundboardSound.
	// This should generally be used opposed to []SoundboardSound.
	SoundboardSoundSlice []*SoundboardSound

	soundboardSoundQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	soundboardSoundType                 = reflect.TypeOf(&SoundboardSound{})
	soundboardSoundMapping              = queries.MakeStructMapping(soundboardSoundType)
	soundboardSoundPrimaryKeyMapping, _ = queries.BindMapping(soundboardSoundType, soundboardSoundMapping, soundboardSoundPrimaryKeyColumns)
	soundboardSoundInsertCacheMut       sync.RWMutex
	soundboardSoundInsertCache          = make(map[string]insertCache)
	soundboardSoundUpdateCacheMut       sync.RWMutex
	soundboardSoundUpdateCache          = make(map[string]updateCache)
	soundboardSoundUpsertCacheMut       sync.RWMutex
	soundboardSoundUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single soundboardSound record from the query using the global executor.
func (q soundboardSoundQuery) OneG(ctx context.Context) (*SoundboardSound, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single soundboardSound record from the query.
func (q soundboardSoundQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SoundboardSound, error) {
	o := &SoundboardSound{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for soundboard_sounds")
	}

	return o, nil
}

// AllG returns all SoundboardSound records from the query using the global executor.
func (q soundboardSoundQuery) AllG(ctx context.Context) (SoundboardSoundSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all SoundboardSound records from the query.
func (q soundboardSoundQuery) All(ctx context.Context, exec boil.ContextExecutor) (SoundboardSoundSlice, error) {
	var o []*SoundboardSound

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SoundboardSound slice")
	}

	return o, nil
}

// CountG returns the count of all SoundboardSound records in the query, and panics on error.
func (q soundboardSoundQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all SoundboardSound records in the query.
func (q soundboardSoundQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count soundboard_sounds rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q soundboardSoundQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q soundboardSoundQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if soundboard_sounds exists")
	}

	return count > 0, nil
}

// SoundboardSounds retrieves all the records using an executor.
func SoundboardSounds(mods ...qm.QueryMod) soundboardSoundQuery {
	mods = append(mods, qm.From("\"soundboard_sounds\""))
	return soundboardSoundQuery{NewQuery(mods...)}
}

// FindSoundboardSoundG retrieves a single record by ID.
func FindSoundboardSoundG(ctx context.Context, iD int, selectCols ...string) (*SoundboardSound, error) {
	return FindSoundboardSound(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindSoundboardSound retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSoundboardSound(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*SoundboardSound, error) {
	soundboardSoundObj := &SoundboardSound{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"soundboard_sounds\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, soundboardSoundObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from soundboard_sounds")
	}

	return soundboardSoundObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *SoundboardSound) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SoundboardSound) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no soundboard_sounds provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(soundboardSoundColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	soundboardSoundInsertCacheMut.RLock()
	cache, cached := soundboardSoundInsertCache[key]
	soundboardSoundInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			soundboardSoundColumns,
			soundboardSoundColumnsWithDefault,
			soundboardSoundColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(soundboardSoundType, soundboardSoundMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(soundboardSoundType, soundboardSoundMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"soundboard_sounds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"soundboard_sounds\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into soundboard_sounds")
	}

	if !cached {
		soundboardSoundInsertCacheMut.Lock()
		soundboardSoundInsertCache[key] = cache
		soundboardSoundInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single SoundboardSound record using the global executor.
// See Update for more documentation.
func (o *SoundboardSound) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the SoundboardSound.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SoundboardSound) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	soundboardSoundUpdateCacheMut.RLock()
	cache, cached := soundboardSoundUpdateCache[key]
	soundboardSoundUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			soundboardSoundColumns,
			soundboardSoundPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update soundboard_sounds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"soundboard_sounds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, soundboardSoundPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(soundboardSoundType, soundboardSoundMapping, append(wl, soundboardSoundPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update soundboard_sounds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for soundboard_sounds")
	}

	if !cached {
		soundboardSoundUpdateCacheMut.Lock()
		soundboardSoundUpdateCache[key] = cache
		soundboardSoundUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q soundboardSoundQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q soundboardSoundQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for soundboard_sounds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for soundboard_sounds")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SoundboardSoundSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SoundboardSoundSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soundboardSoundPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"soundboard_sounds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, soundboardSoundPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in soundboardSound slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all soundboardSound")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *SoundboardSound) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SoundboardSound) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no soundboard_sounds provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(soundboardSoundColumnsWithDefault, o)

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

	soundboardSoundUpsertCacheMut.RLock()
	cache, cached := soundboardSoundUpsertCache[key]
	soundboardSoundUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			soundboardSoundColumns,
			soundboardSoundColumnsWithDefault,
			soundboardSoundColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			soundboardSoundColumns,
			soundboardSoundPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert soundboard_sounds, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(soundboardSoundPrimaryKeyColumns))
			copy(conflict, soundboardSoundPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"soundboard_sounds\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(soundboardSoundType, soundboardSoundMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(soundboardSoundType, soundboardSoundMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert soundboard_sounds")
	}

	if !cached {
		soundboardSoundUpsertCacheMut.Lock()
		soundboardSoundUpsertCache[key] = cache
		soundboardSoundUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single SoundboardSound record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *SoundboardSound) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single SoundboardSound record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SoundboardSound) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SoundboardSound provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), soundboardSoundPrimaryKeyMapping)
	sql := "DELETE FROM \"soundboard_sounds\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from soundboard_sounds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for soundboard_sounds")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q soundboardSoundQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no soundboardSoundQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from soundboard_sounds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for soundboard_sounds")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SoundboardSoundSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SoundboardSoundSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SoundboardSound slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soundboardSoundPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"soundboard_sounds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, soundboardSoundPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from soundboardSound slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for soundboard_sounds")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *SoundboardSound) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no SoundboardSound provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *SoundboardSound) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSoundboardSound(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SoundboardSoundSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty SoundboardSoundSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SoundboardSoundSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SoundboardSoundSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), soundboardSoundPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"soundboard_sounds\".* FROM \"soundboard_sounds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, soundboardSoundPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SoundboardSoundSlice")
	}

	*o = slice

	return nil
}

// SoundboardSoundExistsG checks if the SoundboardSound row exists.
func SoundboardSoundExistsG(ctx context.Context, iD int) (bool, error) {
	return SoundboardSoundExists(ctx, boil.GetContextDB(), iD)
}

// SoundboardSoundExists checks if the SoundboardSound row exists.
func SoundboardSoundExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"soundboard_sounds\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if soundboard_sounds exists")
	}

	return exists, nil
}
