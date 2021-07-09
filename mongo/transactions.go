package mongo

// // Transactions 事务
// type Transactions struct {
// 	ctx  context.Context
// 	name ClientName
// }

// // NewTransactions new transaction
// func NewTransactions(ctx context.Context, name ClientName) *Transactions {
// 	return &Transactions{
// 		ctx:  ctx,
// 		name: name,
// 	}
// }

// // WithTransaction 开启事务
// func (t *Transactions) WithTransaction(f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
// 	client := Client(t.name)
// 	sess, err := client.StartSession()
// 	if err != nil {
// 		return nil, err
// 	}

// 	ccv := cfctx.GetCtxValue(t.ctx).GetCommonValue()
// 	cid := ccv[cfctx.CtxValueCommonKeyCID]
// 	cdb := ccv[cfctx.CtxValueCommonKeyCDB]
// 	// 结束事务
// 	defer sess.EndSession(t.ctx)
// 	r, err := sess.WithTransaction(context.Background(), func(sessCtx mongo.SessionContext) (i interface{}, e error) {
// 		ctx := newCtx(sessCtx, cid, cdb)
// 		return f(ctx)
// 	})

// 	return r, err
// }

// // new ctx
// func newCtx(ctx context.Context, cid, cdb string) context.Context {
// 	ctxMap := make(map[cfctx.CtxValueCommonKey]string)
// 	ctxMap[cfctx.CtxValueCommonKeyCID] = cid
// 	ctxMap[cfctx.CtxValueCommonKeyCDB] = cdb
// 	newCtx, _ := cfctx.SetCtxValue(ctx, cfctx.NewCtxValue(ctxMap, nil, nil))
// 	return newCtx
// }
