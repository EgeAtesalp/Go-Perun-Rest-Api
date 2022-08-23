package ipfs

import (

	"bytes"
	"context"
	"github.com/rs/zerolog/log"
	"crypto/rand"
	"io"
	"github.com/ipfs/go-datastore"
	lvldb "github.com/ipfs/go-ds-leveldb"
	
	"golang.org/x/crypto/nacl/secretbox"
)

type UserInfo struct {

	Alias     string
	Ip_Addr     []byte
	Port   []byte
	Blockchain_Addr []byte

}

func CreateNewDatabase()(*lvldb.Datastore, error){
	
	path := ""

	lvldbds, err := lvldb.NewDatastore("",nil)

	if err != nil {
		log.Error().Msgf("Could not open db connection:%s \n using path: %s \n", err, path)
		return  nil, err
	}

	return lvldbds, nil
}
func AddUserInfo(ui UserInfo, ds *lvldb.Datastore, enckey [32]byte)(error){

	var key datastore.Key = datastore.NewKey(ui.Alias)
	var val = ConvertToByteArray(ui)
	var enval = Encrypt(enckey,val)
	ctx := context.Background()

	err := ds.Put(ctx,key,enval)
	return err
}

func FindUser(alias string, ds *lvldb.Datastore, enckey [32]byte)(UserInfo, error){

	var keyValue datastore.Key = datastore.NewKey(alias)
	ctx := context.Background()
	encval,err := ds.Get(ctx,keyValue)
	if err != nil {
		log.Error().Msgf("Could not find user with alias:%s \n ", alias)
		var nilResponse = UserInfo{"",[]byte(""),[]byte(""),[]byte("")}
		return  nilResponse, err
	}
	val := Decrypt(enckey,encval)
	response := ConvertToUserInfo(val, alias)
	return response, nil
}



func ConvertToUserInfo(val []byte, alias string)(UserInfo){

	s := bytes.Split(val, []byte("/"))
	var ip_address = s[0]
	var port = s[1]
	var block_address = s[2]
	var response = UserInfo{alias, ip_address,port, block_address}
	return response
}

func ConvertToByteArray(ui UserInfo)([]byte){

	var ip_address = ui.Ip_Addr
	var port = ui.Port
	var block_address = ui.Blockchain_Addr
	s := [][]byte{ip_address, port, block_address}
	response := bytes.Join(s, []byte("/"))
	return response
}

func Encrypt(key [32]byte, message []byte)([]byte){
   	
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], message, &nonce, &key)

	return encrypted
}

func Decrypt(key [32]byte, message []byte)([]byte){

	var decryptNonce [24]byte
	copy(decryptNonce[:], message[:24])
	decrypted, ok := secretbox.Open(nil, message[24:], &decryptNonce, &key)
	if !ok {
		panic("decryption error")
	}
	return decrypted

}
