package kafka

// Promise Handlers

func (s *server) handleReadPromise(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleSearchPromises(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleCreatePromise(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleCreatePromiseAndTask(kafkaReq *KafkaRequest) {
	_ = "STUB: not implemented"
	return
}

func (s *server) handleCompletePromise(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleCreateCallback(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleCreateSubscription(kafkaReq *KafkaRequest) {
	_ = "STUB: not implemented"
	return
}

// Schedule Handlers

func (s *server) handleReadSchedule(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleSearchSchedules(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleCreateSchedule(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleDeleteSchedule(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

// Task Handlers

func (s *server) handleClaimTask(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleCompleteTask(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleDropTask(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }

func (s *server) handleHeartbeatTasks(kafkaReq *KafkaRequest) { _ = "STUB: not implemented"; return }
