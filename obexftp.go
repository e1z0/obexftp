package obexftp

/*
#cgo LDFLAGS: -lobexftp
#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include <obexftp/client.h>
// Wrapper for the obexftp_open function
obexftp_client_t* obexftp_open_c(int transport, obex_ctrans_t *ctrans, obexftp_info_cb_t infocb, void *infocb_data) {
    return obexftp_open(transport, ctrans, infocb, infocb_data);
}

// Wrapper for the obexftp_connect_push function
int obexftp_connect_push_c(obexftp_client_t *cli, char *device, int channel) {
    return obexftp_connect_push(cli, device, channel);
}

// Wrapper for the obexftp_put_file function
int obexftp_put_file_c(obexftp_client_t *cli, char *filepath, char *filename) {
    return obexftp_put_file(cli, filepath, filename);
}

// Wrapper for the obexftp_disconnect function
int obexftp_disconnect_c(obexftp_client_t *cli) {
    return obexftp_disconnect(cli);
}

// Wrapper for the obexftp_close function
void obexftp_close_c(obexftp_client_t *cli) {
    obexftp_close(cli);
}
*/
import "C"
import "unsafe"
import "errors"


const (
    VERSION = "0.1"
    AUTHOR = "@e1z0"
)

type ObexFTPClient struct {
    Client *C.obexftp_client_t
}

// Open opens an OBEX FTP session.
func Open() (ObexFTPClient,error) {
    cl := C.obexftp_open(C.int(1), nil, nil, nil)
    if cl == nil {
        return ObexFTPClient{},errors.New("Error opening IrDA device")
    }
    return ObexFTPClient{Client:cl},nil
}

func Connect(cl ObexFTPClient) error {
    ret := C.obexftp_connect_push_c(cl.Client, nil, 0)
    if ret < 0 {
        return errors.New("Error connecting to OBEX remote device")
    }
    return nil
}

func Push(cl ObexFTPClient,filename,filebase string) error {
    filepath := (*C.char)(C.CString(filename))
    basepath := (*C.char)(C.CString(filebase))
    ret := C.obexftp_put_file_c(cl.Client, filepath, basepath)
    C.free(unsafe.Pointer(filepath))
    C.free(unsafe.Pointer(basepath))
    if ret < 0 {
        return errors.New("Unable to upload file")
    } else {
        return nil
    }
}

func Disconnect(cl ObexFTPClient) error {
    ret := C.obexftp_disconnect_c(cl.Client)
    if ret < 0 {
        return errors.New("Error disconnecting from the client")
    } else {
        return nil
    }
}

func Close(cli ObexFTPClient) {
    C.obexftp_close_c(cli.Client)
}

