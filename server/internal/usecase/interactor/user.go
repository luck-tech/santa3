package interactor

import (
	"context"

	"github.com/murasame29/go-httpserver-template/internal/entity"
	"github.com/murasame29/go-httpserver-template/internal/framework/contexts"
	"github.com/murasame29/go-httpserver-template/internal/framework/serrors"
	"github.com/murasame29/go-httpserver-template/internal/usecase/service"
)

type User struct {
	_session        *service.Session
	_user           *service.User
	_skill          *service.Skill
	_usedSkill      *service.UsedSkill
	_wantLearnSkill *service.WantLearnSkill
}

func NewUser(
	session *service.Session,
	user *service.User,
	skill *service.Skill,
	usedSkill *service.UsedSkill,
	wantLearnSkill *service.WantLearnSkill,
) *User {
	return &User{
		_session:        session,
		_user:           user,
		_skill:          skill,
		_usedSkill:      usedSkill,
		_wantLearnSkill: wantLearnSkill,
	}
}

func (u *User) Get(ctx context.Context, id string) (*entity.User, []entity.Skill, []entity.Skill, error) {
	user, found, err := u._user.GetInfo(ctx, id)
	if err != nil {
		return nil, nil, nil, err
	}

	if !found {
		return nil, nil, nil, err
	}

	var usedSkillIDs []int
	var wantLeanSkillIDs []int

	for _, us := range user.UsedSkill {
		usedSkillIDs = append(usedSkillIDs, us.SkillID)
	}

	for _, wls := range user.WantLeanSkills {
		wantLeanSkillIDs = append(wantLeanSkillIDs, wls.SkillID)
	}

	// ここ効率化できるけど面倒だからいったん無視

	usedSkills, err := u._skill.List(ctx, usedSkillIDs)
	if err != nil {
		return nil, nil, nil, err
	}

	wantLeanSkills, err := u._skill.List(ctx, wantLeanSkillIDs)
	if err != nil {
		return nil, nil, nil, err
	}

	return user.User, usedSkills, wantLeanSkills, nil
}

type UpdateUserParam struct {
	UserID          string
	Description     string
	UsedSkills      []string
	WantLearnSkills []string
}

func (u *User) Update(ctx context.Context, param UpdateUserParam) (*entity.User, []entity.Skill, []entity.Skill, error) {
	userID := contexts.GetUserID(ctx)
	user, found, err := u._user.Get(ctx, param.UserID)
	if err != nil {
		return nil, nil, nil, err
	}

	if !found {
		return nil, nil, nil, serrors.ErrUserNotFound
	}

	if user.ID != userID {
		return nil, nil, nil, serrors.ErrPermissionNotFound
	}

	newUser := user

	if param.Description != "" {
		newUser.Description = param.Description
	}

	if err := u._user.Update(ctx, newUser); err != nil {
		return nil, nil, nil, err
	}

	// ここ効率化できるけど面倒だからいったん無視

	if err := u._skill.Upsert(ctx, param.UsedSkills); err != nil {
		return nil, nil, nil, err
	}

	if err := u._skill.Upsert(ctx, param.WantLearnSkills); err != nil {
		return nil, nil, nil, err
	}

	if err := u._usedSkill.UpsertUsedSkill(ctx, param.UserID, param.UsedSkills); err != nil {
		return nil, nil, nil, err
	}

	if err := u._wantLearnSkill.UpsertWantLearnSkill(ctx, param.UserID, param.WantLearnSkills); err != nil {
		return nil, nil, nil, err
	}

	return u.Get(ctx, param.UserID)
}
