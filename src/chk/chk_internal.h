/**
 * (C) Copyright 2022 Intel Corporation.
 *
 * SPDX-License-Identifier: BSD-2-Clause-Patent
 */

/**
 * DAOS global consistency checker RPC Protocol Definitions
 */

#ifndef __CHK_INTERNAL_H__
#define __CHK_INTERNAL_H__

#include <abt.h>
#include <uuid/uuid.h>
#include <daos/rpc.h>
#include <daos/btree.h>
#include <daos/object.h>
#include <daos_srv/pool.h>
#include <daos_srv/daos_chk.h>
#include <daos_srv/daos_engine.h>

#include "chk.pb-c.h"

/*
 * RPC operation codes
 *
 * These are for daos_rpc::dr_opc and DAOS_RPC_OPCODE(opc, ...) rather than
 * crt_req_create(..., opc, ...). See daos/rpc.h.
 */
#define DAOS_CHK_VERSION	1

#define CHK_PROTO_SRV_RPC_LIST									\
	X(CHK_START,										\
		0,	&CQF_chk_start,		ds_chk_start_hdlr,	&chk_start_co_ops),	\
	X(CHK_STOP,										\
		0,	&CQF_chk_stop,		ds_chk_stop_hdlr,	&chk_stop_co_ops),	\
	X(CHK_QUERY,										\
		0,	&CQF_chk_query,		ds_chk_query_hdlr,	&chk_query_co_ops),	\
	X(CHK_MARK,										\
		0,	&CQF_chk_mark,		ds_chk_mark_hdlr,	&chk_mark_co_ops),	\
	X(CHK_ACT,										\
		0,	&CQF_chk_act,		ds_chk_act_hdlr,	&chk_act_co_ops),	\
	X(CHK_REPORT,										\
		0,	&CQF_chk_report,	ds_chk_report_hdlr,	NULL),			\
	X(CHK_REJOIN,										\
		0,	&CQF_chk_rejoin,	ds_chk_rejoin_hdlr,	NULL)

/* Define for RPC enum population below */
#define X(a, b, c, d, e) a

enum chk_rpc_opc {
	CHK_PROTO_SRV_RPC_LIST,
	CHK_PROTO_SRV_RPC_COUNT,
};

#undef X

/* check start in/out */
#define DAOS_ISEQ_CHK_START	/* input fields */				\
	((uint64_t)		(csi_gen)		CRT_VAR)		\
	((uint32_t)		(csi_flags)		CRT_VAR)		\
	((int32_t)		(csi_phase)		CRT_VAR)		\
	((d_rank_t)		(csi_leader_rank)	CRT_VAR)		\
	((uint32_t)		(csi_padding)		CRT_VAR)		\
	((d_rank_t)		(csi_ranks)		CRT_ARRAY)		\
	((struct chk_policy)	(csi_policies)		CRT_ARRAY)		\
	((uuid_t)		(csi_uuids)		CRT_ARRAY)

#define DAOS_OSEQ_CHK_START	/* output fields */				\
	((int32_t)		(cso_status)		CRT_VAR)		\
	((d_rank_t)		(cso_rank)		CRT_VAR)		\
	((uint32_t)		(cso_phase)		CRT_VAR)		\
	((uint32_t)		(cso_padding)		CRT_VAR)		\
	((struct ds_pool_clue)	(cso_clues)		CRT_ARRAY)

CRT_RPC_DECLARE(chk_start, DAOS_ISEQ_CHK_START, DAOS_OSEQ_CHK_START);

/* check stop in/out */
#define DAOS_ISEQ_CHK_STOP	/* input fields */				\
	((uint64_t)		(csi_gen)		CRT_VAR)		\
	((uuid_t)		(csi_uuids)		CRT_ARRAY)

#define DAOS_OSEQ_CHK_STOP	/* output fields */				\
	((int32_t)		(cso_status)		CRT_VAR)		\
	((d_rank_t)		(cso_rank)		CRT_VAR)

CRT_RPC_DECLARE(chk_stop, DAOS_ISEQ_CHK_STOP, DAOS_OSEQ_CHK_STOP);

/* check query in/out */
#define DAOS_ISEQ_CHK_QUERY	/* input fields */				\
	((uint64_t)		(cqi_gen)		CRT_VAR)		\
	((uuid_t)		(cqi_uuids)		CRT_ARRAY)

#define DAOS_OSEQ_CHK_QUERY	/* output fields */				\
	((int32_t)			(cqo_status)	CRT_VAR)		\
	((int32_t)			(cqo_padding)	CRT_VAR)		\
	((struct chk_query_pool_shard)	(cqo_shards)	CRT_ARRAY)

CRT_RPC_DECLARE(chk_query, DAOS_ISEQ_CHK_QUERY, DAOS_OSEQ_CHK_QUERY);

/* check mark in/out */
#define DAOS_ISEQ_CHK_MARK	/* input fields */				\
	((uint64_t)		(cmi_gen)		CRT_VAR)		\
	((d_rank_t)		(cmi_rank)		CRT_VAR)		\
	((uint32_t)		(cmi_version)		CRT_VAR)

#define DAOS_OSEQ_CHK_MARK	/* output fields */				\
	((int32_t)		(cmo_status)		CRT_VAR)

CRT_RPC_DECLARE(chk_mark, DAOS_ISEQ_CHK_MARK, DAOS_OSEQ_CHK_MARK);

/* check act in/out */
#define DAOS_ISEQ_CHK_ACT	/* input fields */				\
	((uint64_t)		(cai_gen)		CRT_VAR)		\
	((uint64_t)		(cai_seq)		CRT_VAR)		\
	((uint32_t)		(cai_cla)		CRT_VAR)		\
	((uint32_t)		(cai_act)		CRT_VAR)		\
	((uint32_t)		(cai_flags)		CRT_VAR)

#define DAOS_OSEQ_CHK_ACT	/* output fields */				\
	((int32_t)		(cao_status)		CRT_VAR)

CRT_RPC_DECLARE(chk_act, DAOS_ISEQ_CHK_ACT, DAOS_OSEQ_CHK_ACT);

/* check report in/out */
#define DAOS_ISEQ_CHK_REPORT	/* input fields */				\
	((uint64_t)		(cri_gen)		CRT_VAR)		\
	((uint32_t)		(cri_ics_class)		CRT_VAR)		\
	((uint32_t)		(cri_ics_action)	CRT_VAR)		\
	((int32_t)		(cri_ics_result)	CRT_VAR)		\
	((d_rank_t)		(cri_rank)		CRT_VAR)		\
	((uint32_t)		(cri_target)		CRT_VAR)		\
	((uint32_t)		(cri_padding)		CRT_VAR)		\
	((uuid_t)		(cri_pool)		CRT_VAR)		\
	((uuid_t)		(cri_cont)		CRT_VAR)		\
	((daos_unit_oid_t)	(cri_obj)		CRT_VAR)		\
	((daos_key_t)		(cri_dkey)		CRT_VAR)		\
	((daos_key_t)		(cri_akey)		CRT_VAR)		\
	((d_string_t)		(cri_msg)		CRT_VAR)		\
	((uint32_t)		(cri_options)		CRT_ARRAY)		\
	((d_sg_list_t)		(cri_details)		CRT_ARRAY)

#define DAOS_OSEQ_CHK_REPORT	/* output fields */				\
	((int32_t)		(cro_status)		CRT_VAR)		\
	((int32_t)		(cro_padding)		CRT_VAR)		\
	((uint64_t)		(cro_seq)		CRT_VAR)

CRT_RPC_DECLARE(chk_report, DAOS_ISEQ_CHK_REPORT, DAOS_OSEQ_CHK_REPORT);

/* check rejoin in/out */
#define DAOS_ISEQ_CHK_REJOIN	/* input fields */				\
	((uint64_t)		(cri_gen)		CRT_VAR)		\
	((d_rank_t)		(cri_rank)		CRT_VAR)		\
	((d_rank_t)		(cri_phase)		CRT_VAR)

#define DAOS_OSEQ_CHK_REJOIN	/* output fields */				\
	((int32_t)		(cro_status)		CRT_VAR)

CRT_RPC_DECLARE(chk_rejoin, DAOS_ISEQ_CHK_REJOIN, DAOS_OSEQ_CHK_REJOIN);

/* dkey for check DB under sys_db */
#define CHK_DB_TABLE		"chk"

/* akey for leader bookmark under CHK_DB_TABLE */
#define CHK_BK_LEADER		"leader"

/* akey for engine bookmark under CHK_DB_TABLE */
#define CHK_BK_ENGINE		"engine"

/* akey for check property under CHK_DB_TABLE */
#define CHK_PROPERTY		"property"

#define CHK_BK_MAGIC_LEADER	0xe6f703da
#define CHK_BK_MAGIC_ENGINE	0xe6f703db
#define CHK_BK_MAGIC_POOL	0xe6f703dc

#define CHK_DUMMY_POOL		"00000000-0000-0000-0000-000020220531"

#define CHK_BTREE_ORDER		16

/*
 * XXX: Please be careful when change CHK__CHECK_INCONSIST_CLASS__CIC_UNKNOWN
 *	to avoid hole is the struct chk_property.
 */
#define CHK_POLICY_MAX		(CHK__CHECK_INCONSIST_CLASS__CIC_UNKNOWN + 1)
#define CHK_POOLS_MAX		(1 << 6)
#define CHK_RANKS_BITMAP_MAX	(1 << 24)

enum chk_act_flags {
	/* The action is applicable to the same kind of inconssitency. */
	CAF_FOR_ALL	= 1,
};

/*
 * Each check instance has a unique leader engine that uses key "chk/leader" under its local
 * sys_db to trace the check instance.
 *
 * For each engine, include the leader engine, there is a system level key "chk/engine" under
 * the engine's local sys_db to trace the check instance on the engine. When server (re)start
 * the check module uses it to determain whether needs to rejoin the check instance.
 *
 * For each pool, there is a key "chk/$pool_uuid" under the engine's local sys_db to trace
 * check process for the pool on related engine.
 */
struct chk_bookmark {
	uint32_t			cb_magic;
	uint32_t			cb_version;
	uint64_t			cb_gen;
	Chk__CheckScanPhase		cb_phase;
	union {
		Chk__CheckInstStatus	cb_ins_status;
		Chk__CheckPoolStatus	cb_pool_status;
	};
	/*
	 * For leader bookmark, it is the inconsistency statistics during the phases range
	 * [CSP_PREPARE, CSP_POOL_LIST] for the whole system. The inconsistency and related
	 * reparation during these phases may be in MS, not related with any engine.
	 *
	 * For pool bookmark, it is the inconsistency statistics during the phases range
	 * [CSP_POOL_MBS, CSP_CONT_CLEANUP] for the pool. The inconsistency and related
	 * reparation during these phases is applied to the pool service leader.
	 */
	struct chk_statistics		cb_statistics;
	struct chk_time			cb_time;
};

/* On each engine (including the leader), there is a key "chk/property" under its local sys_db. */
struct chk_property {
	d_rank_t			cp_leader;
	Chk__CheckFlag			cp_flags;
	Chk__CheckInconsistAction	cp_policies[CHK_POLICY_MAX];
	/*
	 * How many pools will be handled by the check instance. -1 means to handle all pools.
	 * If the specified pools count exceeds CHK_POOLS_MAX, then all pools will be handled.
	 */
	int32_t				cp_pool_nr;
	uuid_t				cp_pools[CHK_POOLS_MAX];
	/*
	 * XXX: Preserve for supporting to continue the check until the specified phase in the
	 *	future. -1 means to check all phases.
	 */
	int32_t				cp_phase;
	/* How many ranks (ever or should) take part in the check instance. */
	uint32_t			cp_rank_nr;
	/* The size of cp_ranks_bitmap in byte. */
	uint32_t			cp_bitmap_sz;
	/*
	 * Bitmap for the ranks that participate in the check instance. It is at most
	 * CHK_RANKS_BITMAP_MAX (16M) bytes, because leader bookmark only exists on the
	 * leader engine, then it is not too much overhead for the whole system.
	 */
	uint8_t				cp_ranks_bitmap[CHK_RANKS_BITMAP_MAX];
};

/*
 * XXX: For each check instance, there are one leader instance and 1 ~ N engine instances.
 *	For each rank, there can be at most one leader instance and one engine instance.
 *
 *	Currently, we do not support to run multiple check instances in the system (even
 *	if they are on different ranks sets) at the same time. If multiple pools need to
 *	be checked, then please either specify their uuids together (or not specify pool
 *	option, then check all pools by default) via single "dmg check" command, or wait
 *	one check instance done and then start next.
 */
struct chk_instance {
	struct chk_bookmark	 ci_bk;
	struct chk_property	 ci_prop;
	/*
	 * For leader, ci_{btr,hdl,list} trace the ranks (engines) that still run check.
	 * For engine, they trace the local pools that are still in checking or pending.
	 */
	union {
		struct btr_root	 ci_rank_btr;
		struct btr_root	 ci_pool_btr;
	};
	union {
		daos_handle_t	 ci_rank_hdl;
		daos_handle_t	 ci_pool_hdl;
	};
	union {
		d_list_t	 ci_rank_list;
		d_list_t	 ci_pool_list;
	};

	struct btr_root		 ci_pending_btr;
	daos_handle_t		 ci_pending_hdl;
	d_list_t		 ci_pending_list;

	/* The slowest phase for the failed pool or rank. */
	uint32_t		 ci_slowest_fail_phase;

	uint32_t		 ci_iv_id;
	struct ds_iv_ns		*ci_iv_ns;
	crt_group_t		*ci_iv_group;

	d_rank_list_t		*ci_ranks;

	ABT_thread		 ci_sched;
	ABT_rwlock		 ci_abt_lock;
	ABT_mutex		 ci_abt_mutex;
	ABT_cond		 ci_abt_cond;

	/* Generator for report event, pending repair actions, and so on. Only for leader. */
	uint64_t		 ci_seq;

	uint32_t		 ci_all_pools:1, /* Check all pools or not. */
				 ci_is_leader:1,
				 ci_sched_running:1,
				 ci_starting:1,
				 ci_stopping:1,
				 ci_started:1,
				 ci_implicated:1;
};

struct chk_iv {
	uint64_t		 ci_gen;
	d_rank_t		 ci_rank;
	uint32_t		 ci_phase;
	uint32_t		 ci_status;
	uint32_t		 ci_to_leader:1;
};

/* Check engine uses it to trace pools. Query logic uses it to organize the result. */
struct chk_pool_shard {
	/* Link into chk_pool_rec::cpr_shard_list. */
	d_list_t		 cps_link;
	d_rank_t		 cps_rank;
	void			*cps_data;
};

/* Check engine uses it to trace pools. Query logic uses it to organize the result. */
struct chk_pool_rec {
	/* Link into chk_instance::ci_pool_list. */
	d_list_t		 cpr_link;
	/* The list of chk_pool_shard. */
	d_list_t		 cpr_shard_list;
	uint32_t		 cpr_shard_nr;
	uint32_t		 cpr_svc_started:1;
	uint32_t		 cpr_phase;
	uuid_t			 cpr_uuid;
	ABT_thread		 cpr_thread;
	struct chk_bookmark	 cpr_bk;
	struct chk_instance	*cpr_ins;
};

struct chk_pending_rec {
	/* Link into chk_instance::ci_pending_list. */
	d_list_t		 cpr_ins_link;
	/* Link into chk_rank_rec::crr_pending_list. */
	d_list_t		 cpr_rank_link;
	uint64_t		 cpr_seq;
	d_rank_t		 cpr_rank;
	uint32_t		 cpr_class;
	uint32_t		 cpr_action;
	uint32_t		 cpr_busy:1,
				 cpr_exiting:1,
				 cpr_on_leader:1;
	ABT_mutex		 cpr_mutex;
	ABT_cond		 cpr_cond;
};

typedef int (*chk_co_rpc_cb_t)(void *args, uint32_t rank, uint32_t phase, int result,
			       void *data, uint32_t nr);

extern struct crt_proto_format	chk_proto_fmt;

extern struct crt_corpc_ops	chk_start_co_ops;
extern struct crt_corpc_ops	chk_stop_co_ops;
extern struct crt_corpc_ops	chk_query_co_ops;
extern struct crt_corpc_ops	chk_mark_co_ops;
extern struct crt_corpc_ops	chk_act_co_ops;

extern btr_ops_t		chk_pool_ops;
extern btr_ops_t		chk_pending_ops;
extern btr_ops_t		chk_rank_ops;

/* chk_common.c */

void chk_ranks_dump(uint32_t rank_nr, d_rank_t *ranks);

void chk_ranks_dump_by_bitmap(uint32_t rank_nr, uint32_t max, uint8_t *bitmap);

void chk_pools_dump(uint32_t pool_nr, uuid_t pools[]);

int chk_bitmap2ranklist(uint32_t rank_nr, d_rank_t max_rank, uint8_t *bitmap,
			d_rank_list_t **rlist);

void chk_stop_sched(struct chk_instance *ins);

int chk_prop_prepare(uint32_t rank_nr, d_rank_t *ranks, uint32_t policy_nr,
		     struct chk_policy **policies, uint32_t pool_nr, uuid_t pools[],
		     uint32_t flags, int phase, d_rank_t leader, bool list_only,
		     struct chk_property *prop, d_rank_list_t **rlist);

int chk_pool_add_shard(daos_handle_t hdl, d_list_t *head, uuid_t uuid, d_rank_t rank,
		       uint32_t phase, struct chk_bookmark *bk, struct chk_instance *ins,
		       uint32_t *shard_nr, void *data);

int chk_pool_del_shard(daos_handle_t hdl, uuid_t pool, d_rank_t rank);

int chk_pending_add(struct chk_instance *ins, d_list_t *rank_head, uint64_t seq,
		    uint32_t rank, uint32_t cla, struct chk_pending_rec **cpr);

int chk_pending_del(struct chk_instance *ins, uint64_t seq, struct chk_pending_rec **cpr);

void chk_pending_destroy(struct chk_pending_rec *cpr);

int chk_ins_init(struct chk_instance *ins);

void chk_ins_fini(struct chk_instance *ins);

/* chk_iv.c */

int chk_iv_update(void *ns, struct chk_iv *iv, uint32_t shortcut, uint32_t sync_mode, bool retry);

int chk_iv_init(void);

int chk_iv_fini(void);

/* chk_leader.c */

int chk_leader_report(uint64_t gen, uint32_t cla, uint32_t act, int32_t result, d_rank_t rank,
		      uint32_t target, uuid_t *pool, uuid_t *cont, daos_unit_oid_t *obj,
		      daos_key_t *dkey, daos_key_t *akey, char *msg, uint32_t option_nr,
		      uint32_t *options, uint32_t detail_nr, d_sg_list_t *details, bool local,
		      uint64_t *seq);

int chk_leader_notify(uint64_t gen, d_rank_t rank, uint32_t phase, uint32_t status);

int chk_leader_rejoin(uint64_t gen, d_rank_t rank, uint32_t phase);

void chk_leader_pause(void);

int chk_leader_init(void);

void chk_leader_fini(void);

/* chk_rpc.c */

int chk_start_remote(d_rank_list_t *rank_list, uint64_t gen, uint32_t rank_nr, d_rank_t *ranks,
		     uint32_t policy_nr, struct chk_policy **policies, uint32_t pool_nr,
		     uuid_t pools[], uint32_t flags, int32_t phase, d_rank_t leader,
		     chk_co_rpc_cb_t start_cb, void *args);

int chk_stop_remote(d_rank_list_t *rank_list, uint64_t gen, uint32_t pool_nr, uuid_t pools[],
		    chk_co_rpc_cb_t stop_cb, void *args);

int chk_query_remote(d_rank_list_t *rank_list, uint64_t gen, uint32_t pool_nr, uuid_t pools[],
		     chk_co_rpc_cb_t query_cb, void *args);

int chk_mark_remote(d_rank_list_t *rank_list, uint64_t gen, d_rank_t rank, uint32_t version);

int chk_act_remote(d_rank_list_t *rank_list, uint64_t gen, uint64_t seq, uint32_t cla,
		   uint32_t act, d_rank_t rank, bool for_all);

int chk_report_remote(d_rank_t leader, uint64_t gen, uint32_t cla, uint32_t act, int32_t result,
		      d_rank_t rank, uint32_t target, char *pool, char *cont, daos_unit_oid_t *obj,
		      daos_key_t *dkey, daos_key_t *akey, char *msg, uint32_t option_nr,
		      uint32_t *options, uint32_t detail_nr, d_sg_list_t *details, uint64_t *seq);

int chk_rejoin_remote(d_rank_t leader, uint64_t gen, d_rank_t rank, uint32_t phase);

/* chk_updcall.c */

int chk_report_upcall(uint64_t gen, uint64_t seq, uint32_t cla, uint32_t act, int32_t result,
		      d_rank_t rank, uint32_t target, uuid_t *pool, uuid_t *cont,
		      daos_unit_oid_t *obj, daos_key_t *dkey, daos_key_t *akey, char *msg,
		      uint32_t option_nr, uint32_t *options, uint32_t detail_nr,
		      d_sg_list_t *details);

/* chk_vos.c */

int chk_bk_fetch_leader(struct chk_bookmark *cbk);

int chk_bk_update_leader(struct chk_bookmark *cbk);

int chk_bk_delete_leader(void);

int chk_bk_fetch_engine(struct chk_bookmark *cbk);

int chk_bk_update_engine(struct chk_bookmark *cbk);

int chk_bk_delete_engine(void);

int chk_bk_fetch_pool(struct chk_bookmark *cbk, uuid_t uuid);

int chk_bk_update_pool(struct chk_bookmark *cbk, uuid_t uuid);

int chk_bk_delete_pool(uuid_t uuid);

int chk_prop_fetch(struct chk_property *cpp);

int chk_prop_update(struct chk_property *cpp);

int chk_traverse_pools(sys_db_trav_cb_t cb, void *args);

void chk_vos_init(void);

void chk_vos_fini(void);

static inline void
chk_query_free(struct chk_query_pool_shard *shards, uint32_t shard_nr)
{
	int	i;

	for (i = 0; i < shard_nr; i++)
		D_FREE(shards[i].cqps_targets);

	D_FREE(shards);
}

#endif /* __CHK_INTERNAL_H__ */
