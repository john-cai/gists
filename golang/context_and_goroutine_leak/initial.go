 ctx, _ := context.WithTimeout(backgroundContext, 10*time.Second)
 doneChan := make(chan *http.Response)
 var err error
 req.WithContext(ctx)
 go func() {
   var err error
   client := http.Client{}
   res, err = client.Do(req)
   if err != nil {
     return 
   }
 
   defer res.Body.Close()
   doneChan <- res
}()

select {
  case res := <-doneChan:
  // parse reponse and save results to database
  case <-ctx.Done():
  if ctx.Err() == context.Canceled{
    log.Info("context was cancelled")
  }
  if ctx.Err() == context.DeadlineExceeded{
    log.Info("request exceeded deadline")
  }
  return
}
