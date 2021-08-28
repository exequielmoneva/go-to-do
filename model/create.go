package model

func CreateTodo() error {
	insertQ, err := con.Query("INSERT INTO TODO($1, $2)", "Exequiel", "Mejorar esta API")
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
