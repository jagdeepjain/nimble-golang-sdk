// Copyright 2020 Hewlett Packard Enterprise Development LP

package service

// AccessControlRecord Service - Manage access control records for volumes.

import (
	"fmt"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/model"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/util"
)

// AccessControlRecordService type 
type AccessControlRecordService struct {
	objectSet *client.AccessControlRecordObjectSet
}

// NewAccessControlRecordService - method to initialize "AccessControlRecordService" 
func NewAccessControlRecordService(gs *NsGroupService) (*AccessControlRecordService) {
	objectSet := gs.client.GetAccessControlRecordObjectSet()
	return &AccessControlRecordService{objectSet: objectSet}
}

// GetAccessControlRecords - method returns a array of pointers of type "AccessControlRecords"
func (svc *AccessControlRecordService) GetAccessControlRecords(params *util.GetParams) ([]*model.AccessControlRecord, error) {
	if params == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",params)
	}
	
	accessControlRecordResp,err := svc.objectSet.GetObjectListFromParams(params)
	if err !=nil {
		return nil,err
	}
	return accessControlRecordResp,nil
}

// CreateAccessControlRecord - method creates a "AccessControlRecord"
func (svc *AccessControlRecordService) CreateAccessControlRecord(obj *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	accessControlRecordResp,err := svc.objectSet.CreateObject(obj)
	if err !=nil {
		return nil,err
	}
	return accessControlRecordResp,nil
}

// UpdateAccessControlRecord - method modifies  the "AccessControlRecord" 
func (svc *AccessControlRecordService) UpdateAccessControlRecord(id string, obj *model.AccessControlRecord) (*model.AccessControlRecord, error) {
	if obj == nil {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",obj)
	}
	
	accessControlRecordResp,err :=svc.objectSet.UpdateObject(id, obj)
	if err !=nil {
		return nil,err
	}
	return accessControlRecordResp,nil
}

// GetAccessControlRecordById - method returns a pointer to "AccessControlRecord"
func (svc *AccessControlRecordService) GetAccessControlRecordById(id string) (*model.AccessControlRecord, error) {
	if len(id) == 0 {
		return nil,fmt.Errorf("error: invalid parameter specified, %v",id)
	}
	
	accessControlRecordResp, err := svc.objectSet.GetObject(id)
	if err != nil {
		return nil,err
	}
	return accessControlRecordResp,nil
}


// DeleteAccessControlRecord - deletes the "AccessControlRecord"
func (svc *AccessControlRecordService) DeleteAccessControlRecord(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("error: invalid parameter specified, %s",id)
	}
	err := svc.objectSet.DeleteObject(id)
	if err != nil {
		return err
	}
	return nil
}