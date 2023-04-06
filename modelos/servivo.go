package modelos

type SerVivo struct{
	Vivo 	bool
}


func (s *SerVivo) EstaVivo() bool{
	s.Vivo = true 
	return s.Vivo
}