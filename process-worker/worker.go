/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package process_worker

import (
	"os"

	"github.com/eolinker/eosc/listener"

	"github.com/eolinker/eosc/traffic"
)

func Process() {

	worker := NewWorker()
	listener.SetTraffic(worker.tf)
	loadPluginEnv()

	worker.wait()
}

type Worker struct {
	tf traffic.ITraffic
}

func (w *Worker) wait() error {
	return nil
}
func NewWorker() *Worker {
	w := &Worker{}
	tf := traffic.NewTraffic()
	tf.Read(os.Stdin)
	w.tf = tf

	return w
}
