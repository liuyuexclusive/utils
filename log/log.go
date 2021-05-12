package log

// // LogToES logrus to elastic
// func LogrusToES(index string) error {
// 	client, err := es.Client()
// 	if err != nil {
// 		logrus.Error(err)
// 		return err
// 	}
// 	hook, err := elogrus.NewElasticHook(client, config.MustGet().ElasticURL, logrus.InfoLevel, index)
// 	if err != nil {
// 		return fmt.Errorf("fail of log to elastic : %w", err)
// 	}
// 	logrus.AddHook(hook)
// 	gin.DefaultWriter = logrus.StandardLogger().Writer()
// 	logrus.Info("seccessfully of log to elastic")
// 	return nil
// }

// func Get() *zap.Logger {
// 	encoderConfig := ecszap.NewDefaultEncoderConfig()
// 	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
// 	logger := zap.New(core, zap.AddCaller())
// 	return logger
// }
