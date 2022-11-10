package graphql

import (
	"context"

	dto "github.com/prometheus/client_model/go"
	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/sensu-go/backend/api"
	"github.com/sensu/sensu-go/backend/store"
	"github.com/stretchr/testify/mock"
)

type MockAssetClient struct {
	mock.Mock
}

func (a *MockAssetClient) ListAssets(ctx context.Context) ([]*corev2.Asset, error) {
	args := a.Called(ctx)
	return args.Get(0).([]*corev2.Asset), args.Error(1)
}

func (a *MockAssetClient) FetchAsset(ctx context.Context, name string) (*corev2.Asset, error) {
	args := a.Called(ctx, name)
	return args.Get(0).(*corev2.Asset), args.Error(1)
}

func (a *MockAssetClient) CreateAsset(ctx context.Context, asset *corev2.Asset) error {
	return a.Called(ctx, asset).Error(0)
}

func (a *MockAssetClient) UpdateAsset(ctx context.Context, asset *corev2.Asset) error {
	return a.Called(ctx, asset).Error(0)
}

type MockCheckClient struct {
	mock.Mock
}

func (c *MockCheckClient) CreateCheck(ctx context.Context, check *corev2.CheckConfig) error {
	return c.Called(ctx, check).Error(0)
}

func (c *MockCheckClient) UpdateCheck(ctx context.Context, check *corev2.CheckConfig) error {
	return c.Called(ctx, check).Error(0)
}

func (c *MockCheckClient) DeleteCheck(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}

func (c *MockCheckClient) ExecuteCheck(ctx context.Context, name string, req *corev2.AdhocRequest) error {
	return c.Called(ctx, name, req).Error(0)
}

func (c *MockCheckClient) FetchCheck(ctx context.Context, name string) (*corev2.CheckConfig, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.CheckConfig), args.Error(1)
}

func (c *MockCheckClient) ListChecks(ctx context.Context) ([]*corev2.CheckConfig, error) {
	args := c.Called(ctx)
	return args.Get(0).([]*corev2.CheckConfig), args.Error(1)
}

type MockSilencedClient struct {
	mock.Mock
}

func (c *MockSilencedClient) UpdateSilenced(ctx context.Context, silenced *corev2.Silenced) error {
	return c.Called(ctx, silenced).Error(0)
}
func (c *MockSilencedClient) GetSilencedByName(ctx context.Context, name string) (*corev2.Silenced, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.Silenced), args.Error(1)
}
func (c *MockSilencedClient) DeleteSilencedByName(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}
func (c *MockSilencedClient) ListSilenced(ctx context.Context) ([]*corev2.Silenced, error) {
	args := c.Called(ctx)
	return args.Get(0).([]*corev2.Silenced), args.Error(1)
}
func (c *MockSilencedClient) GetSilencedByCheckName(ctx context.Context, check string) ([]*corev2.Silenced, error) {
	args := c.Called(ctx, check)
	return args.Get(0).([]*corev2.Silenced), args.Error(1)
}
func (c *MockSilencedClient) GetSilencedBySubscription(ctx context.Context, subs ...string) ([]*corev2.Silenced, error) {
	args := c.Called(ctx, subs)
	return args.Get(0).([]*corev2.Silenced), args.Error(1)
}

type MockHandlerClient struct {
	mock.Mock
}

func (c *MockHandlerClient) ListHandlers(ctx context.Context) ([]*corev2.Handler, error) {
	args := c.Called(ctx)
	return args.Get(0).([]*corev2.Handler), args.Error(1)
}

func (c *MockHandlerClient) FetchHandler(ctx context.Context, name string) (*corev2.Handler, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.Handler), args.Error(1)
}

func (c *MockHandlerClient) CreateHandler(ctx context.Context, handler *corev2.Handler) error {
	return c.Called(ctx, handler).Error(0)
}

func (c *MockHandlerClient) UpdateHandler(ctx context.Context, handler *corev2.Handler) error {
	return c.Called(ctx, handler).Error(0)
}

func (c *MockHandlerClient) DeleteHandler(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}

type MockEntityClient struct {
	mock.Mock
}

func (c *MockEntityClient) DeleteEntity(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}

func (c *MockEntityClient) CreateEntity(ctx context.Context, entity *corev2.Entity) error {
	return c.Called(ctx, entity).Error(0)
}

func (c *MockEntityClient) UpdateEntity(ctx context.Context, entity *corev2.Entity) error {
	return c.Called(ctx, entity).Error(0)
}

func (c *MockEntityClient) FetchEntity(ctx context.Context, name string) (*corev2.Entity, error) {
	args := c.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*corev2.Entity), args.Error(1)
}

func (c *MockEntityClient) ListEntities(ctx context.Context, pred *store.SelectionPredicate) ([]*corev2.Entity, error) {
	args := c.Called(ctx, pred)
	return args.Get(0).([]*corev2.Entity), args.Error(1)
}

type MockEventClient struct {
	mock.Mock
}

func (c *MockEventClient) UpdateEvent(ctx context.Context, event *corev2.Event) error {
	return c.Called(ctx, event).Error(0)
}

func (c *MockEventClient) FetchEvent(ctx context.Context, entity, check string) (*corev2.Event, error) {
	args := c.Called(ctx, entity, check)
	return args.Get(0).(*corev2.Event), args.Error(1)
}

func (c *MockEventClient) DeleteEvent(ctx context.Context, entity, check string) error {
	return c.Called(ctx, entity, check).Error(0)
}

func (c *MockEventClient) ListEvents(ctx context.Context, pred *store.SelectionPredicate) ([]*corev2.Event, error) {
	args := c.Called(ctx, pred)
	return args.Get(0).([]*corev2.Event), args.Error(1)
}

func (c *MockEventClient) ListEventsByEntity(ctx context.Context, entity string, pred *store.SelectionPredicate) ([]*corev2.Event, error) {
	args := c.Called(ctx, entity, pred)
	return args.Get(0).([]*corev2.Event), args.Error(1)
}

func (c *MockEventClient) CountEvents(ctx context.Context, _ *store.SelectionPredicate) (int64, error) {
	args := c.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func (c *MockEventClient) EventStoreSupportsFiltering(ctx context.Context) bool {
	return c.Called(ctx).Get(0).(bool)
}

type MockMutatorClient struct {
	mock.Mock
}

func (c *MockMutatorClient) ListMutators(ctx context.Context) ([]*corev2.Mutator, error) {
	args := c.Called(ctx)
	return args.Get(0).([]*corev2.Mutator), args.Error(1)
}

func (c *MockMutatorClient) FetchMutator(ctx context.Context, name string) (*corev2.Mutator, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.Mutator), args.Error(1)
}

func (c *MockMutatorClient) CreateMutator(ctx context.Context, mutator *corev2.Mutator) error {
	return c.Called(ctx, mutator).Error(0)
}

func (c *MockMutatorClient) UpdateMutator(ctx context.Context, mutator *corev2.Mutator) error {
	return c.Called(ctx, mutator).Error(0)
}

func (c *MockMutatorClient) DeleteMutator(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}

type MockGenericClient struct {
	mock.Mock
}

func (c *MockGenericClient) Create(ctx context.Context, value corev2.Resource) error {
	return c.Called(ctx, value).Error(0)
}

func (c *MockGenericClient) SetTypeMeta(meta corev2.TypeMeta) error {
	return c.Called(meta).Error(0)
}

func (c *MockGenericClient) Update(ctx context.Context, value corev2.Resource) error {
	return c.Called(ctx, value).Error(0)
}

func (c *MockGenericClient) Delete(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}

func (c *MockGenericClient) Get(ctx context.Context, name string, val corev2.Resource) error {
	return c.Called(ctx, name, val).Error(0)
}

func (c *MockGenericClient) List(ctx context.Context, resources interface{}, pred *store.SelectionPredicate) error {
	return c.Called(ctx, resources, pred).Error(0)
}

func (c *MockGenericClient) Authorize(ctx context.Context, verb api.RBACVerb, name string) error {
	return c.Called(ctx, verb, name).Error(0)
}

type MockEventFilterClient struct {
	mock.Mock
}

func (c *MockEventFilterClient) ListEventFilters(ctx context.Context) ([]*corev2.EventFilter, error) {
	args := c.Called(ctx)
	return args.Get(0).([]*corev2.EventFilter), args.Error(1)
}

func (c *MockEventFilterClient) FetchEventFilter(ctx context.Context, name string) (*corev2.EventFilter, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.EventFilter), args.Error(1)
}

func (c *MockEventFilterClient) CreateEventFilter(ctx context.Context, filter *corev2.EventFilter) error {
	return c.Called(ctx, filter).Error(0)
}

func (c *MockEventFilterClient) UpdateEventFilter(ctx context.Context, filter *corev2.EventFilter) error {
	return c.Called(ctx, filter).Error(0)
}

func (c *MockEventFilterClient) DeleteEventFilter(ctx context.Context, name string) error {
	return c.Called(ctx, name).Error(0)
}

type MockNamespaceClient struct {
	mock.Mock
}

func (c *MockNamespaceClient) ListNamespaces(ctx context.Context, pred *store.SelectionPredicate) ([]*corev2.Namespace, error) {
	args := c.Called(ctx, pred)
	return args.Get(0).([]*corev2.Namespace), args.Error(1)
}

func (c *MockNamespaceClient) FetchNamespace(ctx context.Context, name string) (*corev2.Namespace, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.Namespace), args.Error(1)
}

func (c *MockNamespaceClient) CreateNamespace(ctx context.Context, namespace *corev2.Namespace) error {
	return c.Called(ctx, namespace).Error(0)
}

func (c *MockNamespaceClient) UpdateNamespace(ctx context.Context, namespace *corev2.Namespace) error {
	return c.Called(ctx, namespace).Error(0)
}

type MockUserClient struct {
	mock.Mock
}

func (c *MockUserClient) ListUsers(ctx context.Context) ([]*corev2.User, error) {
	args := c.Called(ctx)
	return args.Get(0).([]*corev2.User), args.Error(1)
}

func (c *MockUserClient) FetchUser(ctx context.Context, name string) (*corev2.User, error) {
	args := c.Called(ctx, name)
	return args.Get(0).(*corev2.User), args.Error(1)
}

func (c *MockUserClient) CreateUser(ctx context.Context, user *corev2.User) error {
	return c.Called(ctx, user).Error(0)
}

func (c *MockUserClient) UpdateUser(ctx context.Context, user *corev2.User) error {
	return c.Called(ctx, user).Error(0)
}

type MockMetricGatherer struct {
	mock.Mock
}

func (m *MockMetricGatherer) Gather() ([]*dto.MetricFamily, error) {
	args := m.Called()
	return args.Get(0).([]*dto.MetricFamily), args.Error(1)
}

type MockEtcdHealthController struct {
	mock.Mock
}

func (m *MockEtcdHealthController) GetClusterHealth(ctx context.Context) *corev2.HealthResponse {
	args := m.Called(ctx)
	return args.Get(0).(*corev2.HealthResponse)
}

type MockClusterMetricStore struct {
	mock.Mock
}

func (m *MockClusterMetricStore) EntityCount(ctx context.Context, kind string) (int, error) {
	args := m.Called(ctx, kind)
	return args.Get(0).(int), args.Error(1)
}

type MockRBACClient struct {
	mock.Mock
}

func (m *MockRBACClient) ListRoleBindings(ctx context.Context) ([]*corev2.RoleBinding, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*corev2.RoleBinding), args.Error(1)
}

func (m *MockRBACClient) FetchRoleBinding(ctx context.Context, name string) (*corev2.RoleBinding, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*corev2.RoleBinding), args.Error(1)
}

func (m *MockRBACClient) CreateRoleBinding(ctx context.Context, rb *corev2.RoleBinding) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) UpdateRoleBinding(ctx context.Context, rb *corev2.RoleBinding) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) ListRoles(ctx context.Context) ([]*corev2.Role, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*corev2.Role), args.Error(1)
}

func (m *MockRBACClient) FetchRole(ctx context.Context, name string) (*corev2.Role, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*corev2.Role), args.Error(1)
}

func (m *MockRBACClient) CreateRole(ctx context.Context, rb *corev2.Role) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) UpdateRole(ctx context.Context, rb *corev2.Role) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) ListClusterRoleBindings(ctx context.Context) ([]*corev2.ClusterRoleBinding, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*corev2.ClusterRoleBinding), args.Error(1)
}

func (m *MockRBACClient) FetchClusterRoleBinding(ctx context.Context, name string) (*corev2.ClusterRoleBinding, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*corev2.ClusterRoleBinding), args.Error(1)
}

func (m *MockRBACClient) CreateClusterRoleBinding(ctx context.Context, rb *corev2.ClusterRoleBinding) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) UpdateClusterRoleBinding(ctx context.Context, rb *corev2.ClusterRoleBinding) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) ListClusterRoles(ctx context.Context) ([]*corev2.ClusterRole, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*corev2.ClusterRole), args.Error(1)
}

func (m *MockRBACClient) FetchClusterRole(ctx context.Context, name string) (*corev2.ClusterRole, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(*corev2.ClusterRole), args.Error(1)
}

func (m *MockRBACClient) CreateClusterRole(ctx context.Context, rb *corev2.ClusterRole) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}

func (m *MockRBACClient) UpdateClusterRole(ctx context.Context, rb *corev2.ClusterRole) error {
	args := m.Called(ctx, rb)
	return args.Error(0)
}
