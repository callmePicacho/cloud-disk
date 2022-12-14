// Code generated by goctl. DO NOT EDIT.
package types

type FileUploadChunkCompleteRequest struct {
	Md5        string      `json:"md5"`
	Name       string      `json:"name"`
	Ext        string      `json:"ext"`
	Size       int64       `json:"size"`
	Key        string      `json:"key"`
	UploadId   string      `json:"uploadId"`
	CosObjects []CosObject `json:"cosObjects"` // 上传文件的编号和对应Md5值数组
}

type CosObject struct {
	PartNumber int    `json:"partNumber"`
	Etag       string `json:"etag"`
}

type FileUploadChunkCompleteReply struct {
	Identity string `json:"identity"` // 中心存储池identity
}

type FileUploadChunkRequest struct {
}

type FileUploadChunkReply struct {
	Etag string `json:"etag"` // MD5
}

type FileUploadPrepareRequest struct {
	Md5  string `json:"md5"`  // 上传文件 md5 值
	Name string `json:"name"` // 上传文件名称
}

type FileUploadPrepareReply struct {
	Identity string `json:"identity"` // 共享存储池的唯一标识
	UploadId string `json:"uploadId"`
	Key      string `json:"key"`
}

type RefreshAuthorizationRequest struct {
}

type RefreshAuthorizationReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repositoryIdentity"` // 公共资源池唯一标识
	ParentId           int64  `json:"parentId"`
}

type ShareBasicSaveReply struct {
	Identity string `json:"identity"` // 用户资源池唯一标识
}

type ShareBasicDetailRequest struct {
	Identity string `json:"identity"` // share_basic 的 identity
}

type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repositoryIdentity"` // 公共池中文件的唯一标识
	Name               string `json:"name"`               // 用户资源池中的名称
	Ext                string `json:"ext"`
	Size               int64  `json:"size"`
	Path               string `json:"path"`
}

type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"userRepositoryIdentity"` // 用户池中唯一的标识
	ExpiredTime            int    `json:"expiredTime"`
}

type ShareBasicCreateReply struct {
	Identity string `json:"identity"` // share_basic 的 identity
}

type UserFileMoveRequest struct {
	Identity       string `json:"identity"`
	ParentIdentity string `json:"parentIdentity"`
}

type UserFileMoveReply struct {
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type UserDetailRequest struct {
	Identity string `json:"identity"`
}

type UserDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MailCodeSendRequest struct {
	Email string `json:"email"`
}

type MailCodeSendReply struct {
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type UserRegisterReply struct {
}

type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext  string `json:"ext,optional"`
	Size int64  `json:"size,optional"`
	Path string `json:"path,optional"`
}

type FileUploadReply struct {
	Identity string `json:"identity"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

type UserRepositorySaveRequest struct {
	ParentId           int64  `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}

type UserRepositorySaveReply struct {
}

type UserFileListRequest struct {
	Id   int64 `json:"id,optional"` // 文件层级，等同于 parent_id
	Page int   `json:"page,optional"`
	Size int   `json:"size,optional"`
}

type UserFileListReply struct {
	List  []*UserFile `json:"list"`
	Count int64       `json:"count"`
}

type UserFile struct {
	Id                 int64  `json:"id"`
	Identity           string `json:"identity"`           // 文件记录在个人存储池的唯一标识
	RepositoryIdentity string `json:"repositoryIdentity"` // 文件记录在公共资源池的唯一标识
	Name               string `json:"name"`
	Ext                string `json:"ext"`
	Path               string `json:"path"`
	Size               int64  `json:"size"`
}

type UserFileNameUpdateRequest struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
}

type UserFileNameUpdateReply struct {
}

type UserFolderCreateRequest struct {
	ParentId int64  `json:"parentId"`
	Name     string `json:"name"`
}

type UserFolderCreateReply struct {
	Identity string `json:"identity"`
}

type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}

type UserFileDeleteReply struct {
}
